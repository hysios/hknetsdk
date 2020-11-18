package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/araddon/dateparse"
	"github.com/hysios/hknetsdk/netsdk"
)

var (
	addr     = flag.String("addr", "", "connect address")
	user     = flag.String("user", "", "login username")
	password = flag.String("pass", "", "login password")
	channel  = flag.Int("channel", 1, "channel")
	output   = flag.String("output", "./output.mp4", "下载文件名")
	route    = flag.Int("route", 1, "route")

	startTime = flag.String("start", "", "开始时间")
	duration  = flag.Duration("duration", 5*time.Second, "持续时间")

	// deviceFile = flag.String("input", "", "devices file")
)

func init() {
	flag.Set("addr", "192.168.1.64:8000")
	flag.Set("pass", "adm89679005")
	flag.Set("user", "admin")
}

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		client *netsdk.Client
		err    error
	)

	loc, _ := time.LoadLocation("Asia/Shanghai")

	// var devices []DeviceConfig
	if err = netsdk.Init(); err != nil {
		log.Fatalf("init error %s", err)
	}

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
	log.Printf("serial: %s\n", netsdk.Str(client.DeviceInfo.ST_sSerialNumber[:]))
	log.Printf("channel %d\n", client.DeviceInfo.ST_byStartChan)

	var sTime time.Time
	if len(*startTime) == 0 {
		sTime = time.Now().In(loc).Add(-5 * time.Second)
	} else if sTime, err = dateparse.ParseAny(*startTime); err != nil {
		log.Fatalf("start time format error: %s\n", err)
	}

	defer client.Logout()

	eTime := sTime.Add(*duration)
	var cond = netsdk.NET_DVR_PLAYCOND{
		ST_dwChannel:     uint32(*channel),
		ST_byDownload:    1,
		ST_struStartTime: netsdk.Time2NT(sTime),
		ST_struStopTime:  netsdk.Time2NT(eTime),
	}

	log.Printf("output filenme: %s\n", *output)
	down, err := client.GetFileByTime(*output, &cond)
	if err != nil {
		log.Fatalf("ready download file error %s", err)
	}

	log.Printf("开始下载 %s 到 %s", sTime, eTime)
	down.PlayBack(netsdk.NetDvrPlaystart)
	pos := down.Pos()
	for ; pos < 100; pos = down.Pos() {
		fmt.Printf("downloading %d\r", pos)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("downloading %d\n", pos)
	down.Stop()
	log.Printf("down %#v\n", down)
}
