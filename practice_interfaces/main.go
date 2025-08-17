package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// // this is basic approach
	// bs := make([]byte, 99999) // สร้าง byte slice ที่มีขนาด 99999 ช่อง
	// resp.Body.Read(bs)        // เอา body ที่ได้จาก Resp มาใส่ byte slice ที่เราสร้าง
	// fmt.Println(string(bs))   // ถ้าไม่ใส่ strings จะเป็น 0 0 0 1 0 1 0 ... แต่ถ้าใส่ string จะเป็น <html> ของ google.com

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
