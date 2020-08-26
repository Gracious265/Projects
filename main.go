package main

import (
	"fmt"
	// "math/rand"
	// "github.com/valaparthvi/Projects/numbers"
)


func main() {
	/*fmt.Println(numbers.GetPrimeFactor(rand.Int()))
	ar:= [5]int{1,2,3,4,5}
	br := [5]int{1,2,3,4,5}
	// cr := []int{1,2,3}
	// dr := []int{1,2,3}
	var er []int
	er = ar[:3]
	fmt.Println(ar == br)
	// fmt.Println(cr == dr)
	fmt.Println(er)
	// changing an index value from array will also change it in the slice
	ar[2] = 6
	fmt.Println(er)
	// changing an index value from slice will also change it in the array
	er[2] = 78
	fmt.Println(er, ar)
	fr := [...]int{1,2,3,4,5,6}
	fmt.Println(fr)

	// if len(er) <= len(ar), er will still act as a reference to ar, else not
	er = append(er, 56)
	fmt.Println(er, ar, len(er), cap(er))

	er = append(er, 43, 59, 79, 50)
	fmt.Println(er, ar, len(er), cap(er))

	er = append(er, 100, 101, 90)
	fmt.Println(len(er), cap(er))*/
	gr := []int{1,2,3}
	// var oldCapacity, newCapacity int

	fmt.Println(append(gr, gr...))
}

// TODO: https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94 https://blog.golang.org/slices