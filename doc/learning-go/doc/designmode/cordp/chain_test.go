package cordp

import (
	"fmt"
	"testing"
)

func TestDoc(t *testing.T) {
	Doc()
}

func deal(score, step int, name string) int {
	score += step
	fmt.Printf("%v %v 的值为 %v\n", name, step, score)
	return score
}

func TestChain(t *testing.T) {
	t.Run(fmt.Sprint("TestChain ", 1), func(t *testing.T) {
		// t.Parallel()
		chain := &ChainHandler{}
		chain.AddHandler(func(score int) int {
			return deal(score, 1, t.Name())
		})
		chain.AddHandler(func(score int) int {
			return deal(score, 2, t.Name())
		})
		chain.AddHandler(func(score int) int {
			return deal(score, 3, t.Name())
		})
		chain.Do(1)
	})
	t.Error("t.Error() ---- xixixi")
	fmt.Println("----------")
	t.Run(fmt.Sprint("TestChain ", 2), func(t *testing.T) {
		// t.Parallel()
		chain := &ChainHandler{}
		chain.AddHandler(func(score int) int {
			return deal(score, 1, t.Name())
		})
		chain.AddHandler(func(score int) int {
			return deal(score, 2, t.Name())
		})
		chain.AddHandler(func(score int) int {
			return deal(score, 3, t.Name())
		})
		chain.Do(10)
	})
}
