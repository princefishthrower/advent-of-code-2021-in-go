package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

func main() {
	// Test input
	// fishTimers := []int{3, 4, 3, 1, 2}
	fishCounts := []int{}

	// Real input
	fishTimerStrings := []string{}
	lines := utils.ReadLines("input.txt")
	for _, line := range lines {
		fishTimerStrings = strings.Split(line, ",")
	}

	fishTimers := utils.StringArrayToIntArray(fishTimerStrings)

	// get count of fish at each age
	for age := 0; age < 9; age++ {
		fishCounts = append(fishCounts, utils.CountOccurances(fishTimers, age))
	}

	for i := 0; i < 256; i++ {
		fishCounts = runDay(fishCounts)
		fmt.Println("After Day", i+1, "Sum: ", utils.Sum(fishCounts))
	}

	// solution is simply the sum of fish counts!
	fmt.Println(utils.Sum(fishCounts))
}

// each day, all 0s becomes a 6s and the number of 0 fishes is also added to the 8 count
// then all ages decrease
// here the index IS the count of fish at that age
func runDay(fishCounts []int) []int {
	// store the new fish count
	var countNewFish = fishCounts[0]
	// decrease all
	fishCounts = fishCounts[1:]
	// clear 8 count
	fishCounts = append(fishCounts, 0)
	// now add in additional 6s and new fish
	fishCounts[6] = fishCounts[6] + countNewFish
	fishCounts[8] = fishCounts[8] + countNewFish

	return fishCounts
}
