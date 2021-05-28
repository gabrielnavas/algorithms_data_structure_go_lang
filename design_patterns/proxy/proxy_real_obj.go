package main

import "fmt"

type Proxy struct {
	realObject *RealObject
}

func (p *Proxy) performAction() {
	if p.realObject == nil {
		fmt.Println("realObject is nil.")
		return
	}
	p.realObject.performAction()
}
