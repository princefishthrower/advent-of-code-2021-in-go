package main

import (
	"advent/utils"
	"fmt"
)

func main() {

	// read the lines from the input as string
	stringLines := utils.ReadLines("input.txt")

	// convert the lines from the input to ints
	intLines := utils.StringArrayToIntArray(stringLines)

	// sliding window of 3 over the intLines
	var count = 0
	var prevSum = utils.Sum(intLines[0:3])
	for i := 1; i < len(intLines); i++ {
		var currSum = utils.Sum(intLines[i : i+3])
		if currSum > prevSum {
			count++
		}
		prevSum = currSum
	}

	// print the solution!
	fmt.Println(count)
}
