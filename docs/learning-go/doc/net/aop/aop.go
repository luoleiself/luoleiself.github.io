package aop

import "fmt"

func DoSomething(name string) {
	println("Doing something for", name)
}

func WithLog(fn func(string)) func(string) {
	return func(s string) {
		fmt.Println("Before function:", s)
		fn(s)
		fmt.Println("After function:", s)
	}
}

/* ................................... */

type Service interface {
	Process(data string) error
}
type UserService struct{}

func (u *UserService) Process(data string) error {
	fmt.Println("Processing user:", data)
	return nil
}

type LogProxy struct {
	service Service
}

func (l *LogProxy) Process(data string) error {
	fmt.Println("[LOG] Start:", data)
	err := l.service.Process(data)
	fmt.Println("[LOG] End:", data, "Error:", err)
	return err
}
