package main

import "fmt"

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}
