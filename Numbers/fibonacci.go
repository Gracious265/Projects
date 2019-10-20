package main

import (
	"fmt"
)

func fibonacci (num int) {
	var num0, num1, num2  int
	num0, num1 = 0, 1
	fmt.Println(num0)
	fmt.Println(num1)
	for i:=0 ; i < num -2; i++ {
		num2 = num0 + num1
		num0, num1 = num1, num2
		fmt.Println(num2)
	}
}

func main() {
	var limit int
	fmt.Println("Enter the number: ")
	fmt.Scan(&limit)
	fmt.Println("Starting to print fibonacci series...")
	fibonacci(limit )
}