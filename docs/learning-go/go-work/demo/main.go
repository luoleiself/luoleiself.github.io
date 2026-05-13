package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println("第三方包提供的方法")
	fmt.Println("stringutil.Reverse() ", stringutil.Reverse("hello world"))
	fmt.Println("-----------------------")
	fmt.Println("向第三方包添加的方法")
	fmt.Println("stringutil.ToUpper() ", stringutil.ToUpper("hello world"))
}
