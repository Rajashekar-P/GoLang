package main

import (
	"fmt"
)

func main() {
	fmt.Println("enter a price of product:")
	var price int
	fmt.Scanln(&price)

	fmt.Println("Enter Quantity:")
	var no int
	fmt.Scanln(&no)

	Total := price * no

	fmt.Println("Total is:",Total)
}
