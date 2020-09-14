package number

import (
	"fmt"
	"math/rand"
)

// CoinSimulator : Write some code that simulates flipping a single coin however many times the user decides.
// The code should record the outcomes and count the number of tails and heads.
func CoinSimulator() {
	var (
		number, heads, tails, seed, seedValue, counter int64
		decider                                        int
	)

	fmt.Print("How many times would you like to flip? ")
	fmt.Scan(&number)
	fmt.Print("\nProvide a seed value(int) please: ")
	fmt.Scan(&seed)

	for counter = 0; counter < number; counter++ {
		seedValue = seedValue*seedValue + counter*counter
		rand.Seed(seedValue)
		decider = rand.Intn(2)
		if decider == 0 {
			heads++
		} else {
			tails++
		}
	}

	fmt.Println("Heads:", heads)
	fmt.Println("Tails:", tails)
}
