package main

import "fmt"

func main(){
	c:= add(1,2)
	fmt.Printf("result = %d", c)
}
func add(a, b int) int {
	return a + b
}