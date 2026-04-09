package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Study on time package by go...")
	presentTime := time.Now()
	fmt.Println("Current time :", presentTime)
	fmt.Println("After formatting :\n")
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	dateCreation := time.Date(2027, time.April, 29, 12, 0, 0, 0, time.UTC)
	fmt.Println("\n\nCreated Date is :", dateCreation.Format("02-01-2006 Monday"))
}
