package main

import (
	"fmt"
	"net"
)

func init() {
	fmt.Println("maskFunc.go init()...")
}
func maskFunc() {
	m1 := net.IPv4Mask(255, 255, 0, 0)
	m2 := net.CIDRMask(32, 128)
	fmt.Println("m1 := net.IPv4Mask(255, 255, 0, 0)", m1)
	fmt.Print("m1.Size() ")
	fmt.Println(m1.Size())
	fmt.Println("m1.String() ", m1.String())
	fmt.Println("m2 := net.CIDRMask(32, 128)", m2)
	fmt.Print("m2.Size() ")
	fmt.Println(m2.Size())
	fmt.Println("m2.String() ", m2.String())
	fmt.Println("----------------------------------------")
}
