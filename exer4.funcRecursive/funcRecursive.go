package main

import "fmt"

func main() {
	fmt.Println(fatorial(50))
	fmt.Println(loop(50))
}

func fatorial(x int64) int64 {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loop(x int) (fatorial int) {
	fatorial = 1
	for x > 1 {
		fatorial *= x
		x--
	}
	return
}
