package main

import "fmt"

type Nums [10]int

func (nums Nums) max() (maxValue, index int) {
	index = 0
	maxValue = nums[index]
	for i := 1; i < 10; i++ {
		if nums[i] > maxValue {
			index = i
			maxValue = nums[i]
		}
	}
	return
}

func main() {
	nums := Nums{11, 22, 2, 7, 4, 10, 3, 55, 8, 6}
	maxValue, index := nums.max()
	fmt.Printf("max value %d at index %d.\n", maxValue, index)
}
