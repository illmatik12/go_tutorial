package main

import (
	"fmt"
)

func hello(task chan int) {
	fmt.Println("Hello")
	task <- 0
}

func world() {
	fmt.Println("World!")
}

func main() {
	var task = make(chan int)
	go hello(task)
	<-task
	world()
}
