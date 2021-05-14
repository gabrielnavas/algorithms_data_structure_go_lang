package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func powerSeriesWithReturnNamed(a int) (square, cube int) {
	square = a * a
	cube = a * a * a
	return
}

func main() {
	value := 2
	square, cube := powerSeries(value)
	fmt.Printf("square %d=%d; cube %d=%d\n", value, value, square, cube)

	value2 := 4
	square2, cube2 := powerSeriesWithReturnNamed(value2)
	fmt.Printf("square %d=%d; cube %d=%d\n", value2, value2, square2, cube2)
}
