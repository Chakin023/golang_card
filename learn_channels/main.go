// Program เพื่อ fetching ว่า website นั้น ๆ up หรือไม่ โดยจะ re-fetch เรื่อย ๆ
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//main routine เมื่อ run เสร็จแล้วจะหยุด run ต่อทันทีโดยไม่สนว่า child routine จะ run เสร็จรึเปล่า
	//จึงต้องใช้ channel ในสื่อสารให้ main รอ child ก่อน โดย channel สามารถกำหนด data type สำหรับการรับ-ส่งข้อมูล ซึ่งเป็นชนิดเดียวกัน
	ch := make(chan string)

	for _, link := range links {
		//go routine always in front of called function
		go checkLink(link, ch)
	}

	//any main value to child | ch <- value
	//any var value to ch | var <- ch
	//any function waiting for input | fmt.Println(<-ch)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// }

	for l := range ch {
		go func(link string) {
			//and some pause fetching every 5 seconds by using anonymous function
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}

	fmt.Println(link, "is up!")
	ch <- link
}
