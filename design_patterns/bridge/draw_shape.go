package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct{}

func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing shape")
}
