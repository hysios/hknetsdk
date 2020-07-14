package netsdk

import (
	"bytes"
	"io/ioutil"

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
