package cordp

import (
	"fmt"
	// constraints "golang.org/x/exp/constraints"
)

type Handler func(score int) int

type ChainHandler struct {
	handlers []Handler
}

func (c *ChainHandler) AddHandler(handler Handler) {
	c.handlers = append(c.handlers, handler)
}
func (c *ChainHandler) Do(score int) {
	for _, h := range c.handlers {
		s := h(score)
		fmt.Printf("s 的值为 %v\n", s)
	}
}

func Doc() {
	fmt.Println("CORDP: 责任链模式(Chain Of Responsibility Design Pattern), 类似于洋葱模型")
	fmt.Println("  将请求的发送和接收解耦, 让多个接收对象都有机会消费这个请求,将这些接收对象串成一条链, 并沿着这条链条传递这个请求, 直到链上的某个接收对象能够处理它为止")
}
