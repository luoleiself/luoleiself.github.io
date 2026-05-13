package structdoc

import "strings"

type Student struct {
	name, address string
	age           int
}

type Cat struct {
	Name string
	Age  uint8
	_    struct{}
}

var tab = strings.Repeat(" ", 2)
