package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const newFishLife = 8
const existingNewFishLife = 6

func main() {
	// Test input
	// fishTimers := []int{3, 4, 3, 1, 2}

	// Real input
	lines := utils.ReadLines("input.txt")
	var fishTimers []int
	for _, line := range lines {
		fishTimerStrings := strings.Split(line, ",")
		fishTimers = utils.StringArrayToIntArray(fishTimerStrings)
	}

	for i := 0; i < 80; i++ {
		fishTimers = decreaseByOneAndMultiply(fishTimers)
	}

	// solution is simply the length of fish timers!
	// solution is count of the intersects - print the solution!
	fmt.Println(len(fishTimers))
}

// decreases counter for each fish, adding fish if required
func decreaseByOneAndMultiply(fishTimers []int) []int {
	var newFishTimers []int
	for index := range fishTimers {
		if fishTimers[index] == 0 {
			// for each 0 found after the decrease, reset that fish to 8
			// and append a new 8 fish to the list
			newFishTimers = append(newFishTimers, newFishLife)
			fishTimers[index] = existingNewFishLife
		} else {
			// otherwise normal decrease of 1
			fishTimers[index] = fishTimers[index] - 1
		}
	}
	fishTimers = append(fishTimers, newFishTimers...)
	return fishTimers
}
