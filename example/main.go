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
	"path"
	"strings"
	"time"
	"unsafe"

	"cuelang.org/go/pkg/strconv"
	"github.com/hysios/hknetsdk/netsdk"
	"github.com/kr/pretty"
	"github.com/yudai/pp"
)

var (
	addr     = flag.String("addr", "192.168.1.64:8000", "connect address")
	user     = flag.String("user", "admin", "login username")
	password = flag.String("pass", "adm89679005", "login password")
	// deviceFile = flag.String("input", "", "devices file")
)

func main() {
	flag.Parse()
	var devices []DeviceConfig
	if err := netsdk.Init(); err != nil {
		log.Fatalf("init error %s", err)
	}

	defer netsdk.Cleanup()
	fmt.Println("Hello World!")

	// if deviceFile != nil {
	// 	f, err := os.Open(*deviceFile)
	// 	if err != nil {
	// 		log.Fatalf("open file error %s", *deviceFile)
	// 	}
	// 	defer f.Close()

	// 	devices, err = readDevices(f)
	// 	if err != nil {
	// 		log.Fatalf("read devices error %s", *deviceFile)
	// 	}
	// }

	netsdk.SetExceptionCallback(func(typ int, userData interface{}) bool {
		return true
	}, nil)
	netsdk.SetReconnect(30*time.Second, true)

	client, err := netsdk.Login(*addr, *user, *password)
	log.Printf("%v %s", client, err)
	log.Printf("serial: %s\n", client.DeviceInfo.ST_sSerialNumber[:])
	pp.Printf("deviceinfo %v", client.DeviceInfo)
	var ch = make(chan bool)

	if err := netsdk.SetMessageCallback(nil); err != nil {
		log.Fatalf("set mesasge callback error %s", err)
	}

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
			log.Printf("DeviceID %s\n", string(ainfo.ST_byDeviceID[:]))
			log.Printf("DevInfo IP %s\n", string(ainfo.ST_struDevInfo.ST_struDevIP.ST_sIpV4[:]))
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
				b := C.GoBytes(unsafe.Pointer(pic.ST_pBuffer), C.int(pic.ST_dwDataLen))
				log.Printf("image date %d", len(b))
				imgfile := path.Join(dir, fmt.Sprintf("images/%s-%d.jpg", netsdk.Str(ainfo.ST_byDeviceID[:]), t.UnixNano()))
				err := ioutil.WriteFile(imgfile, b, 0644)
				if err != nil {
					log.Printf("write file error %s", err)
					// return
				}
			}
		}

		t := time.Now()
		jsonfile := path.Join(dir, fmt.Sprintf("images/%d.json", t.UnixNano()))
		f, err := os.Create(jsonfile)
		if err != nil {
			log.Printf("write jsonfile error %s", err)
			return
		}
		defer f.Close()
		enc := json.NewEncoder(f)
		enc.Encode(info)

		log.Printf("info % #v", pretty.Formatter(info))
	}); err != nil {
		log.Fatalf("subcribe client error %s", err)
	}

	defer client.StopSubscribe()
	// msgServ, err := netsdk.StartListen("0.0.0.0", 7200, func(cmd int, pAlarm *netsdk.NET_DVR_ALARMER) {
	// 	log.Printf("cmd %0X\n", cmd)
	// 	log.Printf("SerialNumber %s", string(pAlarm.ST_sSerialNumber[:]))
	// 	log.Printf("DeviceNo %s", string(pAlarm.ST_sDeviceName[:]))
	// 	log.Printf("DeviceIP %s", string(pAlarm.ST_sDeviceIP[:]))
	// 	log.Printf("Port %d", pAlarm.ST_wSocketPort)
	// })
	// if err != nil {
	// 	log.Fatalf("listen message %s", err)
	// }
	// defer msgServ.Stop()
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
