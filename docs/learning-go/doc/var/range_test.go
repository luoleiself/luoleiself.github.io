package main

import (
	"fmt"
	"testing"
)

type P struct {
	Name string
	Age  uint8
}

func TestRangeSlice(t *testing.T) {
	o := []P{{"chain1", 20}, {"chain2", 21}, {"chain3", 22}}
	oPointer := make([]*P, 0, 3)
	t.Run("RangeSlice--1", func(t *testing.T) {
		for _, v := range o {
			oPointer = append(oPointer, &v)
		}
		fmt.Println(oPointer)
		fmt.Println("---------")
	})
	t.Run("RangeSlice--2", func(t *testing.T) {
		for _, v := range oPointer {
			fmt.Println(v)
		}
		fmt.Println("---------")
	})
	t.Run("RangeSlice--3", func(t *testing.T) {
		for _, v := range o {
			v.Age = 18
			fmt.Println(v)
		}
		fmt.Println(o)
		fmt.Println("---------")
	})
	t.Run("RangeSlice--4", func(t *testing.T) {
		for i := range o {
			o[i].Age = 18
		}
		fmt.Println(o)
		fmt.Println("---------")
	})
}
func TestRangeMap(t *testing.T) {
	data := map[string]string{"1": "A", "2": "B", "3": "C", "4": "D"}
	t.Run("RangeMap--1", func(t *testing.T) {
		for k, v := range data {
			data[v] = k
			fmt.Printf("data %v\n", data)
		}
		fmt.Println("---------")
	})
}
