package main

import (
	"advent/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	oxygenGeneratorRatingBinary := getOxygenGeneratorRating()
	co2ScrubberRatingBinary := invertBinaryString(oxygenGeneratorRatingBinary)
	fmt.Println(oxygenGeneratorRatingBinary)
	fmt.Println(co2ScrubberRatingBinary)

	// convert both gammaRate and epsilonRate to
	gammaRate, err := strconv.ParseInt(oxygenGeneratorRatingBinary, 2, 0)
	if err != nil {
		log.Fatalf("Cannot convert to int: %s", oxygenGeneratorRatingBinary)
	}
	epsilonRate, err := strconv.ParseInt(co2ScrubberRatingBinary, 2, 0)
	if err != nil {
		log.Fatalf("Cannot convert to int: %s", co2ScrubberRatingBinary)
	}

	// solution is the depth multiplied by the position
	solution := gammaRate * epsilonRate

	// print the solution!
	fmt.Println(solution)
}

func getOxygenGeneratorRating() string {
	stringLines := utils.ReadLines("input.txt")
	// firstItemLength := len(stringLines[0])
	length := len(stringLines)

	// how to pass this 12 in dynamically?
	zeroCounts := [12]int{}
	for _, stringLine := range stringLines {
		chars := strings.Split(stringLine, "")
		for index, char := range chars {
			if char == "0" {
				zeroCounts[index] += 1
			}
		}
	}

	// now that we have zero counts at each position, build the gammaRate string
	gammaRate := ""
	for _, zeroCount := range zeroCounts {
		// if the zero count is greater than 50%, it is the most common bit
		if zeroCount == (length / 2) {
			gammaRate += "0"
		} else if zeroCount > (length / 2) {
			gammaRate += "0"
		} else {
			gammaRate += "1"
		}
	}
	return gammaRate
}

// invert 1s to 0s, and likewise 0s to 1s in a binary string
func invertBinaryString(binaryString string) string {
	chars := strings.Split(binaryString, "")
	inverse := ""
	for _, char := range chars {
		if char == "0" {
			inverse += "1"
		} else {
			inverse += "0"
		}
	}
	return inverse
}
