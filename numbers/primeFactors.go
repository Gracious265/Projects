package numbers


import (
	"fmt"
	"math"
)

func getPrimeFactor() {
	var number int
	fmt.Println("Enter a number to find it's prime factors: ")
	fmt.Scan(&number)
	for number%2 == 0 {
		fmt.Println(2)
		number = number/2
	}
	for i:=3; i<=int(math.Sqrt(float64(number)))+1;{
		for number % i==0 {
			fmt.Println(number)
			number=number/i
		}
	}
	if number>2{
		fmt.Println(number)
	}

}