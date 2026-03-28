package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the review for the recently purchased product : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your review is : ", input)
}
