package number

import (
	"fmt"
)

// GetCostOfTiles : Calculate the total cost of tile it would take to cover a floor plan of width and height, using a cost entered by the user.
func GetCostOfTiles() {
	var fW, fH, tW, tH, cost, totalTiles, totalCost float64
	fmt.Print("Enter the floor Width and Height: ")
	fmt.Scan(&fW, &fH)
	fmt.Print("Enter the tile Width and Height: ")
	fmt.Scan(&tW, &tH)
	fmt.Print("Enter the cost per tile: ")
	fmt.Scan(&cost)
	totalTiles = (fW * fH) / (tW * tH)
	totalCost = totalTiles * cost
	fmt.Println(totalCost)
}
