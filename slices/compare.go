package main

import "fmt"

func main() {
	a := []byte("string")
	b := []byte("strinh")

	switch Compare(a, b) {
	case -1:
		fmt.Println("\nResult:\n\tA less than B")
	case 0:
		fmt.Println("\nResult:\n\tA equals B")
	case 1:
		fmt.Println("\nResult:\n\tA greater than B")
	}
}

func Compare(a, b []byte) int {

	for i := 0; i < len(a) && i < len(b); i++ {
		fmt.Printf("[%v] A:%v | B:%v\n", i, a[i], b[i])
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
