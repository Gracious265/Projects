package main

import (
	// "github.com/valaparthvi/Projects/number"
	"fmt"
	"github.com/codingsince1985/geo-golang/openstreetmap"

	"github.com/codingsince1985/geo-golang"
)

const (
	addr         = "Melbourne VIC"
	lat, lng     = -37.813611, 144.963056
	radius       = 50
	zoom         = 18
	addrFR       = "Champs de Mars Paris"
	latFR, lngFR = 48.854395, 2.304770
)

func main() {
	exampleGeocoder()
}
func exampleGeocoder() {
	fmt.Println("OpenStreetMap")
	try(openstreetmap.Geocoder())
}
func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := geocoder.ReverseGeocode(lat, lng)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", lat, lng, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Print("\n")
}

func tryOnlyFRData(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addrFR)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addrFR, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := geocoder.ReverseGeocode(latFR, lngFR)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", latFR, lngFR, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Print("\n")
}
