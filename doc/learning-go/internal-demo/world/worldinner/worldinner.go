package worldinner

import (
	"fmt"
	"internal-demo/world/internal/internala"
	"internal-demo/world/internal/internalb"
)

func WSay() {
	fmt.Println("worldinner package, WSay func...")
	
	fmt.Println("calling internala(b) package in worldinner package...")
	internala.ISay()
	internalb.ISay()
}
