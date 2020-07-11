package main

import "fmt"

func main() {

	a := x()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := x()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func x() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
