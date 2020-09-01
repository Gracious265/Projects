package number

import (
	"fmt"
	"strconv"
	"strings"
)

// UnitConverter : Converts various units between one another.
// The user enters the type of unit being entered, the type of unit they want to convert to and then the value.
// The program will then make the conversion.
func UnitConverter() {
	var unitType string
	fmt.Print("Enter the type of conversion (temp/mass): ")
	fmt.Scan(&unitType)
	switch unitType {
	case "temp":
		var tempF string
		var temp float64
		fmt.Print("Enter the temperature (append -C for Celcius, -F for Farenheit) : ")
		fmt.Scan(&tempF)
		a := strings.Split(tempF, "-")
		temp, _ = strconv.ParseFloat(a[1], 64)
		switch a[1] {
		case "C":
			// (0°C × 9/5) + 32 = 32°F
			fTemp := (temp * 9 / 5) + 32
			fmt.Printf("Temperature in Fahrenheit is: %.2f F", fTemp)
		case "F":
			// (0°F − 32) × 5/9
			cTemp := (temp - 32) * 5 / 9
			fmt.Printf("Temperature in Celcius is: %.2f C", cTemp)
		}
	}
}
