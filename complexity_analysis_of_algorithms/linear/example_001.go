package main

import "fmt"

func main() {
	var nums [10]int
	for i := 0; i < 10; i++ {
		nums[i] = i * 200
		fmt.Println(nums[i])
	}
	fmt.Println("this linear complexity is O(n), where n=10")
}
