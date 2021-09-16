package main

func main() {
	var ch chan int
	close(ch)
	ch <- 1
	//runtime.closechan()
	//runtime.chansend1()
}
