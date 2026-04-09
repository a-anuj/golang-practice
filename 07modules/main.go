package main

import (
	"bufio"
	"fmt"
	"os"
)

func greeting(name string) string {
	stringToReturn := "Greetings, " + name
	return stringToReturn
}

func main() {
	fmt.Println("Enter your name : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(greeting(input))
}
