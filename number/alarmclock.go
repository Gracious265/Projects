package number

import (
	"fmt"
	"time"
	"math"
	
	// "io"
)

// AlarmClock : A simple clock where it plays a sound after X number of minutes/seconds or at a particular time.
func AlarmClock() {
	var minutes int64
	var duration time.Duration
	fmt.Print("Enter the number of minutes after which to ring the alarm: ")
	fmt.Scan(&minutes)
	minutes *= int64(math.Pow(10, 9))
	// TODO: Fix this.
	duration = time.Duration{minutes}
	fmt.Printf("%v", minutes)
	time.Sleep(duration)
	fmt.Println("\nHey there!")
}
