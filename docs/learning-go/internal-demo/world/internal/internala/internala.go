package internala

import (
	"fmt"
	"internal-demo/world/internal/internalb"
)

func ISay() {
	fmt.Println("internala package, ISay func...")

	fmt.Println("calling internalb package in internala package...")
	internalb.ISay()
}
