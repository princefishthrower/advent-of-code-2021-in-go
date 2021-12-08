package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

type FuelScore struct {
	Position      int
	TotalFuelCost int
}

func main() {
	// Test input
	// crabHorizontalPosition := []int{16,1,2,0,4,2,7,1,2,14}

	// Real input
	lines := utils.ReadLines("input.txt")
	var crabHorizontalPositions []int
	for _, line := range lines {
		crabHorizontalPositionStrings := strings.Split(line, ",")
		crabHorizontalPositions = utils.StringArrayToIntArray(crabHorizontalPositionStrings)
	}

	minPosition := 0
	maxPosition := utils.Max(crabHorizontalPositions)

	// calculate fuel cost at all positions
	var fuelScores []FuelScore
	for i := minPosition; i < maxPosition; i++ {
		totalFuelCost := totalFuelCostForPosition(crabHorizontalPositions, i)
		fuelScore := FuelScore{i, totalFuelCost}
		fuelScores = append(fuelScores, fuelScore)
	}

	// solution is the FuelScore with the minimum cost!
	minFuelScore := minFuelScore(fuelScores)
	fmt.Println(minFuelScore)
}

func totalFuelCostForPosition(crabHorizontalPositions []int, position int) int {
	var fuelCosts []int
	for _, crabHorizontalPosition := range crabHorizontalPositions {
		// only difference from part 1 is here: the fuelCost calculation is a bit more complex
		fuelCostSimple := utils.Abs(crabHorizontalPosition - position)
		fuelCostComplex := infiniteSeries(fuelCostSimple)
		fuelCosts = append(fuelCosts, fuelCostComplex)
	}
	return utils.Sum(fuelCosts)
}

func minFuelScore(fuelScores []FuelScore) FuelScore {
	var minIndex int = 0
	var min = fuelScores[0]
	for index, value := range fuelScores {
		if min.TotalFuelCost > value.TotalFuelCost {
			min = value
			minIndex = index
		}
	}
	return fuelScores[minIndex]
}

func infiniteSeries(number int) int {
	return (number * (number + 1)) / 2
}
