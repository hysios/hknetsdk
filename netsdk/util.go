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
