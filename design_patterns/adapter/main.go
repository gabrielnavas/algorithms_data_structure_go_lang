package main

import (
	"fmt"
)

func main() {
	var processor IProcess = AdapterIntToString{}
	num := processor.process()
	fmt.Printf("The number on str is: %s\n", num)
}
