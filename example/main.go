package main

import "C"

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"cuelang.org/go/pkg/strconv"
	"github.com/hysios/hknetsdk/netsdk"
	"github.com/kr/pretty"
	"github.com/yudai/pp"
)

var (
	addr     = flag.String("addr", "", "connect address")
	user     = flag.String("user", "", "login username")
	password = flag.String("pass", "", "login password")
	channel  = flag.Int("channel", 1, "channel")
	route    = flag.Int("route", 1, "route")

	// deviceFile = flag.String("input", "", "devices file")
)

func init() {
	flag.Set("addr", "192.168.1.64:8000")
	flag.Set("pass", "adm89679005")
	flag.Set("user", "admin")
}

var safeStateMap = map[netsdk.BYTE]string{
	0: "未知",
	1: "系安全带",
	2: "未系安全带",
}

var callStateMap = map[netsdk.BYTE]string{
	0: "未知",
	1: "没打电话",
	2: "打电话",
}

var detectModeMap = map[netsdk.BYTE]string{
	1: "地感触发",
	2: "视频触发",
	3: "多帧识别",
	4: "雷达触发",
}

func main() {
	flag.Parse()

	var (
		client *netsdk.Client
		err    error
	)

	// var devices []DeviceConfig
	if err = netsdk.Init(); err != nil {
		log.Fatalf("init error %s", err)
	}

	sigs := make(chan os.Signal, 1)
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)

		client.StopSubscribe()
		netsdk.Cleanup()

		os.Exit(-1)
	}()

	defer netsdk.Cleanup()

	netsdk.SetExceptionCallback(func(typ int, userData interface{}) bool {
		log.Printf("exception %d %s %v", typ, netsdk.ExceptionString(typ), userData)
		log.Printf("Last Error %s", netsdk.Err(netsdk.GetLastErrorN()))

		return true
	}, nil)
	netsdk.SetReconnect(30*time.Second, true)
	netsdk.SetRecvTimeOut(15 * time.Second)

	client, err = netsdk.Login(*addr, *user, *password)
	log.Printf("%v %s", client, err)
	log.Printf("serial: %s\n", client.DeviceInfo.ST_sSerialNumber[:])
	pp.Printf("deviceinfo %v", client.DeviceInfo)
	var ch = make(chan bool)
	cruise, err := client.GetPTZCruise(*channel, *route)
	if err != nil {
		log.Printf("get ptz cruise error %s", err)
	} else {
		pp.Printf("cruse settings %v", cruise)
	}

	cfg, err := client.DVRConfig(netsdk.NetDvrGetDevicecfgV40, 0)
	if err != nil {
		log.Printf("get dvr config error %s", err)
	} else {
		if devicecfg, ok := cfg.(*netsdk.NET_DVR_DEVICECFG_V40); ok {
			pp.Printf("Dvr Config %s", devicecfg)
			log.Printf("DVRName %s\n", netsdk.Str(devicecfg.ST_sDVRName[:]))
			log.Printf("SerialNumber %s\n", netsdk.Str(devicecfg.ST_sSerialNumber[:]))
			log.Printf("DevTypeName %s\n", netsdk.Str(devicecfg.ST_byDevTypeName[:]))
		}
	}

	cfg, err = client.DVRConfig(netsdk.NetDvrGetTrackParamcfg, 1)
	if err != nil {
		log.Printf("get dvr TrackParamcfg error %s", err)
	} else {
		if devicecfg, ok := cfg.(*netsdk.NET_DVR_TRACK_PARAMCFG); ok {
			pp.Printf("Dvr TrackParamcfg %s", devicecfg)
			devicecfg.ST_wTrackHoldTime = 60
			devicecfg.ST_byTrackMode = 2
			log.Println(client.SetDVRConfig(netsdk.NetDvrSetTrackParamcfg, 1, devicecfg))
		}

	}

	go func() {
		time.Sleep(3 * time.Second)
		plateResult, err := client.Capture(&netsdk.NET_DVR_MANUALSNAP{
			ST_byLaneNo:  0,
			ST_byChannel: 0,
		})

		if err != nil {
			log.Printf("dvr Capture error %s", err)
		} else {
			pp.Printf("plateResult %s", plateResult)
		}

		var snapcfg = netsdk.NET_DVR_SNAPCFG{
			ST_dwSize:        uint32(unsafe.Sizeof(netsdk.NET_DVR_SNAPCFG{})),
			ST_bySnapTimes:   1,
			ST_wSnapWaitTime: 1000,
			ST_wIntervalTime: [4]uint16{100, 0, 0, 0},
		}
		err = client.ContinuousShoot(&snapcfg)
		if err != nil {
			log.Printf("ContinuousShoot error %s", err)
		} else {
			pp.Printf("plateResult %s", plateResult)
		}
	}()

	// out, err := client.DeviceAbility(netsdk.DADeviceSofthardware, "")
	// if err != nil {
	// 	log.Printf("DeviceAbility: DeviceSofthardware error %s", err)
	// } else {
	// 	log.Printf("DeviceAbility: %s\n", string(out))
	// }
	// config, err := client.ConfigFile()
	// if err != nil {
	// 	log.Fatalf("get config error %s", err)
	// }
	// log.Printf("%s\n", string(config))

	if err := netsdk.SetMessageCallback(nil); err != nil {
		log.Fatalf("set mesasge callback error %s", err)
	}

	var prefix = "unknown"
	// ret, err := client.GetDVRConfig(netsdk.NetItcGetTriggercfg, 0)
	// if err != nil {
	// 	log.Fatalf("get dvr config %s", err)
	// }

	// cfg := ret.(*netsdk.NET_ITC_TRIGGERCFG)

	// log.Printf("ITC_POST_HVT_PARAM_V50 % #v", pretty.Formatter(cfg.ST_struTriggerParam.NET_ITC_POST_HVT_PARAM_V50()))

	if err := client.Subscribe(0, func(cmd netsdk.CommAlarm, cli *netsdk.Client, alarm *netsdk.NET_DVR_ALARMER, info interface{}) {
		log.Printf("cmd 0x%0x", cmd)
		log.Printf("SerialNumber %s", string(alarm.ST_sSerialNumber[:]))
		log.Printf("DeviceNo %s", string(alarm.ST_sDeviceName[:]))
		log.Printf("DeviceIP %s", string(alarm.ST_sDeviceIP[:]))
		log.Printf("Port %d", alarm.ST_wSocketPort)
		dir, _ := os.Getwd()
		log.Printf("current path %s", dir)
		switch cmd {
		case netsdk.CommAlarmTfs:
			log.Println("交通违法信息上报")
			ainfo, ok := info.(*netsdk.NET_DVR_TFS_ALARM)
			if !ok {
				log.Printf("invalid type ")
				return
			}
			prefix = "TFS"
			log.Printf("DeviceID %s\n", string(ainfo.ST_byDeviceID[:]))
			log.Printf("DevInfo IP %s\n", string(ainfo.ST_struDevInfo.ST_struDevIP.ST_sIpV4[:]))
			b := netsdk.Bytes(ainfo.ST_struPlateInfo.ST_sLicense[:])
			b, _ = netsdk.GbkToUtf8(b)
			log.Printf("PlateInfo %s\n", string(b))
			log.Printf("Plate Believe %s\n", string(ainfo.ST_struPlateInfo.ST_byBelieve[:]))

			picInfos := ainfo.ST_struPicInfo
			for _, pic := range picInfos {
				if pic.ST_dwDataLen == 0 {
					continue
				}
				t := time.Now()
				n := strings.Replace(t.Format("20060102150405.999999"), ".", "", 1)
				b := C.GoBytes(unsafe.Pointer(pic.ST_pBuffer), C.int(pic.ST_dwDataLen))
				log.Printf("image type %d image data len %d", pic.ST_byType, len(b))
				imgfile := path.Join(dir, fmt.Sprintf("images/%s-%s.jpg", netsdk.Str(ainfo.ST_byDeviceID[:]), n))
				err := ioutil.WriteFile(imgfile, b, 0644)
				if err != nil {
					log.Printf("write file error %s", err)
					// return
				}
			}
		case netsdk.CommAlarmAidV41:
			log.Println("交通事件信息")
			ainfo, ok := info.(*netsdk.NET_DVR_AID_INFO)
			if !ok {
				log.Printf("invalid type ")
				return
			}
			prefix = "AID"
			log.Printf("ainfo %#v", ainfo)
		case netsdk.CommUploadPlateResult:
			log.Println("卡口车辆上报")
			t := time.Now()
			ainfo, ok := info.(*netsdk.NET_DVR_PLATE_RESULT)
			if !ok {
				log.Printf("invalid type ")
				return
			}
			prefix = "UploadPlate"
			log.Printf("info % #v", info)
			b := netsdk.Bytes(ainfo.ST_struPlateInfo.ST_sLicense[:])
			b, _ = netsdk.GbkToUtf8(b)
			log.Printf("AbsTime %s\n", string(ainfo.ST_byAbsTime[:]))
			log.Printf("PlateInfo %s\n", string(b))

			if ainfo.ST_pBuffer1 != nil {
				b := C.GoBytes(unsafe.Pointer(ainfo.ST_pBuffer1), C.int(ainfo.ST_dwPicLen))
				log.Printf("image date %d", len(b))
				imgfile := path.Join(dir, fmt.Sprintf("images/%d.jpg", t.UnixNano()))
				err := ioutil.WriteFile(imgfile, b, 0644)
				if err != nil {
					log.Printf("write file error %s", err)
					return
				}
			}

			if ainfo.ST_pBuffer2 != nil {
				b := C.GoBytes(unsafe.Pointer(ainfo.ST_pBuffer2), C.int(ainfo.ST_dwPicPlateLen))
				log.Printf("image date %d", len(b))
				imgfile := path.Join(dir, fmt.Sprintf("images/plate-%d.jpg", t.UnixNano()))
				err := ioutil.WriteFile(imgfile, b, 0644)
				if err != nil {
					log.Printf("write file error %s", err)
					return
				}
			}
		case netsdk.CommItsPlateResult:
			log.Println("卡口二次识别车辆上报")
			ainfo, ok := info.(*netsdk.NET_ITS_PLATE_RESULT)
			if !ok {
				log.Printf("invalid type ")
				return
			}
			prefix = "ITS"
			log.Printf("info % #v", info)
			b := netsdk.Bytes(ainfo.ST_struPlateInfo.ST_sLicense[:])
			b, _ = netsdk.GbkToUtf8(b)
			// log.Printf("AbsTime %s\n", string(ainfo.ST_byAbsTime[:]))
			log.Printf("PlateInfo %s\n", string(b))
			log.Printf("触发车道 %d", ainfo.ST_byDriveChan)
			log.Printf("驾驶员系安全带 %s", safeStateMap[ainfo.ST_byPilotSafebelt])
			log.Printf("副驾驶系安全带 %s", safeStateMap[ainfo.ST_byCopilotSafebelt])
			log.Printf("驾驶员打电话 %s", callStateMap[ainfo.ST_byPilotCall])
			log.Printf("速度上限 %d", ainfo.ST_wSpeedLimit)
			log.Printf("违章类型 %d", ainfo.ST_wIllegalType)
			log.Printf("违章子类型 %s", netsdk.Str(ainfo.ST_byIllegalSubType[:]))
			// log.Printf("驾驶员系安全带 %d",	ainfo.ST_byPilotSafebelt)
			// log.Printf("驾驶员系安全带 %d",	ainfo.ST_byPilotSafebelt)
			log.Printf("车辆类型 %d", ainfo.ST_byVehicleType)
			log.Printf("危险车辆类型 %d", ainfo.ST_byDetSceneID)
			log.Printf("通道报警类型 %d", ainfo.ST_byVehiclePositionControl)
			log.Printf("检测方式 %s", detectModeMap[ainfo.ST_byDetectType])
			log.Printf("自定义违章类型 %d", ainfo.ST_dwCustomIllegalType)
			log.Printf("违章信息类型 %d", ainfo.ST_byIllegalFromatType)

			if ainfo.ST_byIllegalFromatType == 1 {
				b := C.GoBytes(unsafe.Pointer(ainfo.ST_pIllegalInfoBuf), 64)
				s := netsdk.Str(b)
				log.Printf("违法类型编号 %s", s)
			}
			log.Printf("车俩违法类型 %d", ainfo.ST_struVehicleInfo.ST_byIllegalType)

			// ST_byPilotSafebelt   BYTE // 0-表示未知,1-系安全带,2-不系安全带
			// ST_byCopilotSafebelt BYTE // 0-表示未知,1-系安全带,2-不系安全带
			// ST_byPilotSunVisor   BYTE // 0-表示未知,1-不打开遮阳板,2-打开遮阳板
			// ST_byCopilotSunVisor BYTE // 0-表示未知, 1-不打开遮阳板,2-打开遮阳板
			// ST_byPilotCall       BYTE //  0-表示未知, 1-不打电话,2-打电话
			picInfos := ainfo.ST_struPicInfo
			for _, pic := range picInfos {
				if pic.ST_dwDataLen == 0 {
					continue
				}
				t := time.Now()
				n := strings.Replace(t.Format("20060102150405.999999"), ".", "", 1)

				b := C.GoBytes(unsafe.Pointer(pic.ST_pBuffer), C.int(pic.ST_dwDataLen))
				log.Printf("image type %d image data len %d", pic.ST_byType, len(b))
				imgfile := path.Join(dir, fmt.Sprintf("images/%s-%s.jpg", netsdk.Str(ainfo.ST_byDeviceID[:]), n))
				err := ioutil.WriteFile(imgfile, b, 0644)
				if err != nil {
					log.Printf("write file error %s", err)
					// return
				}
			}
		case netsdk.CommGisinfoUpload:
			log.Printf("镜头变化")
			ainfo, ok := info.(*netsdk.NET_DVR_GIS_UPLOADINFO)
			if !ok {
				log.Printf("invalid type ")
				return
			}
			prefix = "GIS"
			log.Printf("PTS pos %f", ainfo.ST_struPtzPos.ST_fPanPos)
			log.Printf("Tilt Pos %f", ainfo.ST_struPtzPos.ST_fTiltPos)
			log.Printf("Zoom pos %f", ainfo.ST_struPtzPos.ST_fZoomPos)
		default:
			log.Printf("info % #v", pretty.Formatter(info))
		}

		t := time.Now()
		n := strings.Replace(t.Format("20060102150405.999999"), ".", "", 1)
		jsonfile := path.Join(dir, fmt.Sprintf("images/%s-%s.json", prefix, n))
		f, err := os.Create(jsonfile)
		if err != nil {
			log.Printf("write jsonfile error %s", err)
			return
		}
		defer f.Close()
		enc := json.NewEncoder(f)
		enc.Encode(info)

	}); err != nil {
		log.Fatalf("subcribe client error %s", err)
	}

	defer client.StopSubscribe()

	<-ch
}

type DeviceConfig struct {
	IP   string
	Port int64
	User string
	Pass string
}

func readDevices(rd io.ReadCloser) ([]DeviceConfig, error) {
	var (
		s       = bufio.NewScanner(rd)
		devices = make([]DeviceConfig, 0)
	)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		var (
			line     = s.Text()
			sections = strings.Fields(line)
			devCfg   DeviceConfig
		)
		if len(sections) >= 4 {
			devCfg.IP = sections[0]
			devCfg.Port, _ = strconv.ParseInt(sections[0], 10, 32)
			devCfg.User = sections[3]
			devCfg.Pass = sections[4]
			devices = append(devices, devCfg)
		}
	}

	return devices, nil
}
