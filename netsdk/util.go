package netsdk

import (
	"bytes"
	"io/ioutil"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Str(b []byte) string {
	n := bytes.IndexByte(b, 0x00)
	if n >= 0 {
		return string(b[:n])
	}
	return string(b)
}

func Bytes(b []byte) []byte {
	n := bytes.IndexByte(b, 0x00)
	if n >= 0 {
		return b[:n]
	}
	return b
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func NTEx2Time(nt NET_DVR_TIME_EX) time.Time {
	return time.Date(
		int(nt.ST_wYear),
		time.Month(nt.ST_byMonth),
		int(nt.ST_byDay),
		int(nt.ST_byHour),
		int(nt.ST_byMinute),
		int(nt.ST_bySecond),
		0,
		time.Local,
	)
}

func NTExV2Time(nt NET_DVR_TIME_V30) time.Time {

	return time.Date(
		int(nt.ST_wYear),
		time.Month(nt.ST_byMonth),
		int(nt.ST_byDay),
		int(nt.ST_byHour),
		int(nt.ST_byMinute),
		int(nt.ST_bySecond),
		int(nt.ST_wMilliSec)*1000000,
		time.Local,
	)
}

func NT2Time(nt NET_DVR_TIME) time.Time {
	return time.Date(
		int(nt.ST_dwYear),
		time.Month(nt.ST_dwMonth),
		int(nt.ST_dwDay),
		int(nt.ST_dwHour),
		int(nt.ST_dwMinute),
		int(nt.ST_dwSecond),
		0,
		time.Local,
	)
}

func Time2NT(t time.Time) NET_DVR_TIME {
	return NET_DVR_TIME{
		ST_dwYear:   uint32(t.Year()),
		ST_dwMonth:  uint32(t.Month()),
		ST_dwDay:    uint32(t.Day()),
		ST_dwHour:   uint32(t.Hour()),
		ST_dwMinute: uint32(t.Minute()),
		ST_dwSecond: uint32(t.Second()),
	}
}
