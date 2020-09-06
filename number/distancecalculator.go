package number

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"math"
)

const (
	earthRadius float64 = 6371
)

type city struct {
	name                string
	latitude, longitude float64
}

func getDistance(source, destination city) float64 {
	// Haversine formula:	a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
	// c = 2 ⋅ atan2( √a, √(1−a) )
	// d = R ⋅ c
	// 	where	φ is latitude, λ is longitude, R is earth’s radius (mean radius = 6,371km);
	// note that angles need to be in radians to pass to trig functions!
	var distance float64
	fmt.Println(source, destination)
	a := math.Pow(math.Sin((math.Abs(destination.latitude-source.latitude)/2)*math.Pi/180), 2) + (math.Cos(source.latitude*math.Pi/180)*math.Cos(destination.latitude*math.Pi/180))*math.Pow(math.Sin((math.Abs(destination.longitude-source.longitude)/2)*math.Pi/180), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance = earthRadius * c
	return distance
}

// DistanceBetweenTwoCities : Calculates the distance between two cities and allows the user to specify a unit of distance.
// This program may require finding coordinates for the cities like latitude and longitude.
func DistanceBetweenTwoCities() {
	var (
		source, destination city
		geocoder            geo.Geocoder
	)
	fmt.Print("Enter source city : ")
	fmt.Scan(&source.name)
	fmt.Print("Enter destination city : ")
	fmt.Scan(&destination.name)
	geocoder = openstreetmap.Geocoder()
	location, _ := geocoder.Geocode(source.name)
	source.latitude, source.longitude = location.Lat, location.Lng
	location, _ = geocoder.Geocode(destination.name)
	destination.latitude, destination.longitude = location.Lat, location.Lng
	fmt.Println(getDistance(source, destination))
}

// TODO:  https://www.movable-type.co.uk/scripts/latlong.html
// https://github.com/dnf0/pyprojects/blob/master/citydistance.py
// Calculate distance between 2 cities by checking the program in main.go
// https://www.geeksforgeeks.org/program-distance-two-points-earth/
// https://github.com/codingsince1985/geo-golang/blob/master/openstreetmap/geocoder.go
