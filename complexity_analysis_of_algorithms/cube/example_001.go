package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				fmt.Println("the loop is ", i, j, k, "is ", 1)
			}
		}
	}
}
