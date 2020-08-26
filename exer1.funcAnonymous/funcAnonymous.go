// Working with Anonymous function in Golang
package main

import (
	"fmt"
)

func main() {

	// # Anonymous function
	// calling anonymous function immediately after function definition
	a := 10

	func(i int) {
		fmt.Println("[Anonymous function]:", i*100)
	}(a)
	// calling anonymous function assigning to a variable
	b := func(i int) int {
		return i * 200
	}
	fmt.Println("[Anonymous function assigned to a variable]:", b(10))

	// # Return function from another function
	c := returnFunc()
	fmt.Println("[Returned funtion]:", c(10))
	// Calling the returned function directly
	fmt.Println("[Returned funtion]:", returnFunc()(10))

}

func returnFunc() func(i int) int {
	return func(i int) int {
		return i * 300
	}
}
