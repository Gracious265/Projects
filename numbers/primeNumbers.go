package numbers

import (
	"fmt"
	"math"
	"os"
)

// PrimeNumbers : Have the program find prime numbers until the user chooses to stop asking for the next one.
func PrimeNumbers() {
	var i, j, counter, limit int
	var answer string
	answer = "Y"
	limit = 20
	for i = 2; i <= limit; i++ {
		counter = 0
		for j = 2; j <= int(math.Sqrt(float64(i)))+1; j++ {
			if i%j == 0 {
				counter++
			}
		}
		if counter == 0 {
			fmt.Println(i)
			fmt.Print("Press Y to get the next prime number, else N: ")
			fmt.Scan(&answer)
			if answer == "N" {
				fmt.Println("Exiting the program ...")
				os.Exit(0)
			}
		}
	}
}
