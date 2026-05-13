package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("--------------------------")
	fmt.Println("func New(text string) error 根据给定的文本返回一个错误类型的值")
	fmt.Println("func Is(err, target error) bool 判断 err 是否是 Target 的实例")

	var err = fmt.Errorf("fmt.Errorf(format string, a ...interface{}) error // generate error %s", "hello world")
	fmt.Printf("err 的类型为 %T 值为 %v\n", err, err)
	fmt.Println("--------------------------")

	var err2 = errors.New("errors.New(text string) error // generate error")
	fmt.Printf("err2 的类型为 %T 值为 %v\n", err2, err2)

	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path: ", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("-----------------------")

	var myErr = errors.New("my error")
	myErr2 := fmt.Errorf("my error: %w", myErr)
	fmt.Printf("使用 == 结果 %t\n", myErr == myErr2)                                // false
	fmt.Printf("使用 errors.Is(myErr2, myErr) 结果 %t\n", errors.Is(myErr2, myErr)) // true
	fmt.Println("-----------------------")

	err3 := errors.New("error3")
	err4 := fmt.Errorf("error4: %w", err3)
	fmt.Println(err4)
	fmt.Println(errors.Unwrap(err4))
}

// type BaseError struct {
// 	code uint8
// 	msg  string
// 	data interface{}
// }

// func (be *BaseError) Error() string {
// 	return ""
// }
