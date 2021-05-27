package main

func main() {
	process := &ProcessClass{}

	decorator := &ProcessDecorator{}
	// decorator dont have process
	decorator.process()

	// now have
	decorator.processInstance = process
	decorator.process()
}
