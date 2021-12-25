// HEAD 메서드로 헤더가져오기
package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	log.Println(resp.Status)     // 문자열로 "200 OK"
	log.Println(resp.StatusCode) // 수치로 200
	log.Panicln(resp.Header)
}
