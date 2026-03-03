package main

import (
	"fmt"
	"internal-demo/hello"
	"internal-demo/world"
	// "internal-demo/world/internal/internala"
	// "internal-demo/world/internal/internalb"
)

func main() {
	fmt.Println("internal-demo module main package main func...")

	hello.Say()

	world.Say()

	// internala.ISay()
	// internalb.ISay()
	// $ go run .
	// package internal-demo
	// 				hello.go:7:2: use of internal package internal-demo/world/internal/internala not allowed
	// package internal-demo
	// 				hello.go:8:2: use of internal package internal-demo/world/internal/internalb not allowed
	// fmt.Println("calling internala(b) package fun in main package...")
}
