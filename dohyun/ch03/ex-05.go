// POST 메서드로 임의의 바디 전송
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("ex-05.go")
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		// 전송 실패
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
