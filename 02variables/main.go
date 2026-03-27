package main

import "fmt"

const name string = "Hello"

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

	var smallFloat float64 = 255.98732648736422736478253467253445372654762354
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of the type : %T\n", smallFloat)

	var aliasVariable int
	fmt.Println(aliasVariable)
	fmt.Printf("Variable is of the type : %T\n", aliasVariable)

	// No var keyword
	fmt.Println("---------------------------------------------------")

	temp := 7
	fmt.Printf("My lucky number is %d\n", temp)

	fmt.Printf("My name is %s\n", name)
}
