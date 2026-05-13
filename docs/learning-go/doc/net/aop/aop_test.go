package aop

import "testing"

func TestFuncAOP(t *testing.T) {
	t.Run("函数直接调用", func(t *testing.T) {
		DoSomething("test")
	})
	t.Run("函数封装 AOP", func(t *testing.T) {
		WithLog(DoSomething)("test")
	})
}

func TestProxyAOP(t *testing.T) {
	t.Run("Proxy AOP", func(t *testing.T) {
		var s Service = &UserService{}
		s = &LogProxy{service: s}
		s.Process("test")
	})
}
