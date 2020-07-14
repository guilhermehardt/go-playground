// Simple example of Callback in GO
package main

import "fmt"

func main() {

	years := []int{}
	for i := 2010; i <= 2050; i++ {
		years = append(years, i)
	}
	fmt.Printf("Between %v and %v the World Cup should happen:\n", years[0], years[len(years)-1])
	c, y := worldCup(count, years...)
	fmt.Printf("\t- Times: %v \n\t- Range: %v\n", c, y)

}

func count(i ...int) int {
	return len(i)
}

func worldCup(f func(i ...int) int, years ...int) (count int, worldCupYears []int) {

	for i, y := range years {
		if i%4 == 0 {
			worldCupYears = append(worldCupYears, y)
		}
	}
	count = f(worldCupYears...)
	return
}
