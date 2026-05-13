package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "test.txt"))
	if err != nil {
		t.Fatalf("Open file failure, %v\n", err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	// buf.Size() // 返回以字节为单位基础缓冲区的大小
	// buf.Buffered() //  返回当前缓冲区可读取的字节数

	for {
		line, err := buf.ReadString('\n') // ReadLine 是低水平方法, 使用 ReadString('\n') 代替
		if err != nil && err != io.EOF {
			t.Fatalf("Read file failure, %v\n", err)
		}
		if err == io.EOF {
			fmt.Printf("result = %v\n", line)
			fmt.Println("Read Complete!!!")
			break
		}
		fmt.Println("result =", line)
	}
}
