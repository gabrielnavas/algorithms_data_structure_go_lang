package main

func main() {
	var proxy = &Proxy{&RealObject{}}
	proxy.performAction()
}
