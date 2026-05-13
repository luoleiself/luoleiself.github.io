package world

import (
	"fmt"

	"internal-demo/world/internal/internala"
	"internal-demo/world/internal/internalb"
	"internal-demo/world/worldinner"
)

func Say() {
	fmt.Println("internal-demo module world package, Say func....")

	fmt.Println("calling worldinner package func in world package...")
	worldinner.WSay()

	fmt.Println("calling internala(b) package func in world package...")
	internala.ISay()
	internalb.ISay()
}
