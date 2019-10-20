package main

import (
	"math"
	"fmt"
)

func main() {
	var limit string
	fmt.Println("Enter a number: ")
	fmt.Scan(&limit)
	fmt.Println(limit)
	pi := "%." + limit + "f"
	fmt.Println(fmt.Sprintf(pi, math.Pi))
}