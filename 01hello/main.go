package main

import "fmt"

func helper() {
	fmt.Println("Helper is running!!")
}

func main() {
	go helper()
	fmt.Println("Hello from the World of go ")
}
