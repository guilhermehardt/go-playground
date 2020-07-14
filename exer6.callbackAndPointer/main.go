package main

import (
	"fmt"
)

func main() {
	a := appendMulti3()

	for x := 1; x < 10; x++ {
		fmt.Println(a(&[]int{x}))
	}

}

func multi3(v, l int) (s int) {
	return v * l
}

func appendMulti3() func(x *[]int) []int {
	valor := []int{}
	return func(x *[]int) []int {
		for _, v := range *x {
			valor = append(valor, multi3(v, len(valor)))
		}
		return valor
	}
}
