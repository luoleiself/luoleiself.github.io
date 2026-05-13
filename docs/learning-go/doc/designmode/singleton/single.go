package singleton

import (
	"fmt"
	"sync"
)

type singleTon struct {
	Name, Address string
	Age           uint8
}

func (s *singleTon) SayHello() {
	fmt.Printf("%s is %d year old, from %s\n", s.Name, s.Age, s.Address)
}

var instance *singleTon
var once sync.Once

func newSingleTon() *singleTon {
	// sync.Once 确保只执行一次
	once.Do(func() {
		instance = &singleTon{}
	})
	return instance
}
