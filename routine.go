package main

import (
	"time"
	"fmt"
)

func main(){
	fmt.Println("main start.")
	go f1()
	go f2()
	fmt.Println("main Process.")
	time.Sleep(time.Second*2)
	fmt.Println("main End.")

}
func f1(){
	fmt.Println("F1 Start.")
	time.Sleep(time.Second)
	fmt.Println("F1 End.")
}
func f2(){
	fmt.Println("F2 Start.")
}