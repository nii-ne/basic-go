package main

import (
	"time"
	"fmt"
)

var c = make(chan string)

func pong() {
	var str string
	for{
		str = <-c
		fmt.Println(str)
		if str == "ping" {
			fmt.Println("pong")
		} else {
			fmt.Println("Error")
		}
	}
}

func main() {
	go pong()
	fmt.Println("main start")
	c <- "ping"
	fmt.Println("main end")
	time.Sleep(time.Second)
}
