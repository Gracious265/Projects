package number

import (
	"fmt"
	"time"
)

// AlarmClock : A simple clock where it plays a sound after X number of minutes/seconds or at a particular time.
func AlarmClock() {
	var seconds int64
	fmt.Print("Enter the number of seconds after which to ring the alarm: ")
	fmt.Scan(&seconds)
	time.Sleep(time.Duration(seconds)*time.Second)
	fmt.Println("\nHey there!")
// Check this : https://github.com/hajimehoshi/go-mp3/blob/master/example/main.go
}
