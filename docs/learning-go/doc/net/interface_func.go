package main

import (
	"fmt"
	"net"
)

func init() {
	fmt.Println("interfaceFunc.go init()1...")
}
func init() {
	fmt.Println("interfaceFunc.go init()2...")
}
func interfaceFunc() {
	resByIndex, _ := net.InterfaceByIndex(14)
	fmt.Printf("resByIndex, _ := net.InterfaceByIndex(14)\n %+v\n", resByIndex)
	fmt.Println("----------------------------------------")
	fmt.Println(tab, "func InterfaceAddrs()([]Addrs, error) // 返回系统的网络接口的地址列表")
	res1, _ := net.InterfaceAddrs()
	fmt.Println("res1, _ := net.InterfaceAddrs()\nlen(res1)", len(res1))
	fmt.Println("res1 = ", res1)
	fmt.Println("res1[0].Network()", res1[0].Network(), "res1[0].String()", res1[0].String())
	fmt.Println("res1[1].Network()", res1[1].Network(), "res1[1].String()", res1[1].String())
	fmt.Println("----------------------------------------")
	res2, _ := net.Interfaces()
	res3, _ := res2[0].Addrs()
	res4, _ := res2[0].MulticastAddrs()
	fmt.Println("res2, _ := net.Interfaces()\nlen(res2)", len(res2))
	fmt.Println("res2[0].Addrs()", res3)
	fmt.Println("res2[0].MulticastAddrs()", res4)
	res5, _ := res2[1].Addrs()
	res6, _ := res2[1].MulticastAddrs()
	fmt.Println("res2[1].Addrs()", res5)
	fmt.Println("res2[1].MulticastAddrs()", res6)
	fmt.Println("---------------------")

	fmt.Println(tab, "func Interfaces() ([]Interface, error) // 接口返回一个当前系统的网络接口信息结构体组成的切片, 可以遍历切片获取每个接口信息")
	fmt.Println(tab, "Interface 结构体包含当前网络接口的索引,最大传输单元，接口名，硬件地址，接口的属性信息")
	fmt.Println(`
	for i := 0; i < len(res2); i++ {
		fmt.Printf("res2[%\d].Index 为 %\v, res2[%\d].MTU 为 %\v, res2[%\d].Name 为 %\v, res2[%\d].HardwareAddr 为 %\v, res2[%\d].Flags 为 %\v \n", i, res2[i].Index, i, res2[i].MTU, i, res2[i].Name, i, res2[i].HardwareAddr, i, res2[i].Flags)
	}
	`)
	for i := 0; i < len(res2); i++ {
		fmt.Printf("res2[%d].Index 为 %v, res2[%d].MTU 为 %v, res2[%d].Name 为 %v, res2[%d].HardwareAddr 为 %v, res2[%d].Flags 为 %v \n", i, res2[i].Index, i, res2[i].MTU, i, res2[i].Name, i, res2[i].HardwareAddr, i, res2[i].Flags)
	}
	fmt.Println("----------------------------------------")
}
