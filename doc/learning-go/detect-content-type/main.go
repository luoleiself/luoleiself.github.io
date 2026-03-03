package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("./demo.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// http.DetectContentType 判断文件类型
	// 获取缓冲区的前 n 个字节
	contentType := http.DetectContentType(buffer[:n])
	// demo.jpg -> image/jpeg
	// demo.mp4 -> video/mp4
	// demo.pdf -> application/octet-stream
	// demo.zip -> application/octet-stream
	fmt.Println(contentType)
}
