package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hello world")
	new()
}

func new() {
	arr := [5]int{1, 2, 3, 4, 5}
	for key := range arr {
		fmt.Println(key)
	}
}
