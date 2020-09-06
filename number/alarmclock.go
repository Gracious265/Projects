package number

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"log"
	"os"
	"time"
)

func playMusic() error {
	f, err := os.Open("classic.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

// AlarmClock : A simple clock where it plays a sound after X number of minutes/seconds or at a particular time.
func AlarmClock() {
	var seconds int64
	fmt.Print("Enter the number of seconds after which to ring the alarm: ")
	fmt.Scan(&seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	if err := playMusic(); err != nil {
		log.Fatal(err)
	}
	// Check this : https://github.com/hajimehoshi/go-mp3/blob/master/example/main.go
}
