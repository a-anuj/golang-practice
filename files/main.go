package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Gonna write 'Hi' to the file....")
	file, err := os.Create("test.txt")
	handleError(err)
	content := "Hii"
	length, err := io.WriteString(file, content)
	handleError(err)

	fmt.Println("Total Length written is :", length)
	defer file.Close()
	readFile("test.txt")
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	handleError(err)
	fmt.Println("The data present inside the file is :", string(data))
}
