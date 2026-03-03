package hello

import (
	"fmt"

	"internal-demo/world"
	// "internal-demo/world/internal/internala"
	"internal-demo/world/worldinner"
)

func Say() {
	fmt.Println("internal-demo module hello package, Say func...")

	fmt.Println("calling world package func in hello package...")
	world.Say()

	fmt.Println("calling world package / worldinner package func in hello package...")
	worldinner.WSay()

	// internala.ISay()
	// $ go run .
	// package internal-demo
	//       imports internal-demo/hello
	//       hello\hello.go:7:2: use of internal package internal-demo/world/internal/internala not allowed
	// fmt.Println("calling internala package func in hello package...")
}
