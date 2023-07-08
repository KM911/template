package main

import (
	"abtest/request"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/KM911/oslib/adt"
	"io"
)

func UnGzip(data string) string {
	buf := bytes.NewBuffer([]byte(data))
	zr, err := gzip.NewReader(buf)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer zr.Close()

	var out bytes.Buffer
	io.Copy(&out, zr)
	return out.String()
}
func Gzip(data string) []byte {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Write([]byte(data))
	zw.Close()
	return buf.Bytes()

}

func main() {
	defer adt.TimerStart().End()
	headers := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "*/*",
		"Content-Encoding": "gzip",
	}
	client := request.NewPostClient("http://192.168.0.102:3000/test",
		headers, `{"name":"dzg","age":18}`)
	fmt.Println("this res", client.Send())
}
