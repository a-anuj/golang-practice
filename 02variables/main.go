package main

import "fmt"

func main() {
	var username string = "Anuj"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T\n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of the type : %T\n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of the type : %T\n", smallVal)
}
