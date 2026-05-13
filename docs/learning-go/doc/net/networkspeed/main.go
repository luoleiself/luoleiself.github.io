package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func measureDownloadSpeed(url string) (float64, error) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	bytesDownload, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		return 0, err
	}
	duration := time.Since(start).Seconds() // 从 start 到现在经过的时间, 计算过去的时间

	return float64(bytesDownload) / duration / (1024 * 1024), nil
}

// func measureUploadSpeed(url string, dataSize int) (float64, error) {
// 	data := make([]byte, dataSize)
// 	start := time.Now()

// 	response, err := http.Post(url, "application/octet-stream", io.NopCloser(bytes.NewBuffer(data)))
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer response.Body.Close()

// 	duration := time.Since(start).Seconds()
// 	return float64(dataSize) / duration / (1024 * 1024), nil
// }

func main() {
	url := "http://gips3.baidu.com/it/u=1821127123,1149655687&fm=3028&app=3028&f=JPEG&fmt=auto?w=720&h=1280"
	speed, err := measureDownloadSpeed(url)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf("Download Speed: %.2f MB/s\n", speed)

	// uploadUrl := "https://www.baidu.com/upload"
	// uploadSpeed, err := measureUploadSpeed(uploadUrl, 10*1024*1024)
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// fmt.Printf("Upload speed: %.2fMB/s\n", uploadSpeed)
}
