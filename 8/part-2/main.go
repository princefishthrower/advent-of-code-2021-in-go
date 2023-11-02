package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

// note that these constants represent the segment letters ONLY
// the segments themselves are in alphabetical order but otherwise represent no
// sort of spatial order

// var zero = []string{"a", "b", "c", "e", "f", "g"}
var one = []string{"c", "f"}

// var two = []string{"a", "c", "d", "e", "g"}
// var three = []string{"a", "c", "d", "f", "g"}
var four = []string{"b", "c", "d", "f"}

// var five = []string{"a", "b", "d", "f", "g"}
// var six = []string{"a", "b", "d", "e", "f", "g"}
var seven = []string{"a", "c", "f"}
var eight = []string{"a", "b", "c", "d", "e", "f", "g"}

// var nine = []string{"a", "b", "c", "d", "f", "g"}

// counts of how many segments are needed to produce the respective numbers
const oneLength = 2
const fourLength = 4
const sevenLength = 3
const eightLength = 7

func main() {
	lines := utils.ReadLines("input.txt")

	sum := 0
	// for each note, get the segment output part of the note
	// then count the 1s,4s,7s, and 8s in the outputs
	for _, line := range lines {
		noteLine := strings.Split(line, " | ")
		inputs := strings.Split(noteLine[0], " ")
		outputs := strings.Split(noteLine[1], " ")
		number := getNumbers(outputs)
		sum += number
	}

	// solution is the sums of the outputs!
	fmt.Println(sum)
}

func getNumbers(outputs []string) int {
	count := 0
	for _, output := range outputs {
		outputLetters := strings.Split(output, "")
		number = determineNumber(outputLetters)
		numbers = append(numbers, number)

	}
	return count
}

func determineNumber(outputLettersCount int) int {
	switch outputLettersCount {
	case 2:
		return 1
	case 3:
		// could be
	case 4:
		return 4
	case sevenLength:
		return 7
	case eightLength:
		return 8
	case lengthSix:
		// could be a
	}
}
