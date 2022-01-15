// 멀티파트 폼을 POST로 전송하기
package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "dohyun")

	fileWriter, err := writer.CreateFormFile("thumbnail", "dohyun.png")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("dohyun.png")
	if err != nil {
		// 파일 읽기 실패
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		// 전송 실패
		panic(err)
	}
	log.Println("Status: ", resp.Status)

}
