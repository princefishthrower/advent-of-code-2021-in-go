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

	counts := 0
	// for each note, get the segment output part of the note
	// then count the 1s,4s,7s, and 8s in the outputs
	for _, line := range lines {
		noteLine := strings.Split(line, " | ")
		// inputs := strings.Split(noteLine[0], " ")
		outputs := strings.Split(noteLine[1], " ")
		count := countOnesFoursSevensAndEights(outputs)
		counts += count
	}

	// solution is the counts of 1, 4, 7, and 8 in all the outputs
	fmt.Println(counts)
}

func countOnesFoursSevensAndEights(outputs []string) int {
	count := 0
	for _, output := range outputs {
		outputLetters := strings.Split(output, "")
		if isOneFourSevenOrEight(len(outputLetters)) {
			count++
		}
	}
	return count
}

func isOneFourSevenOrEight(outputLettersCount int) bool {
	switch outputLettersCount {
	case oneLength:
		return true
	case fourLength:
		return true
	case sevenLength:
		return true
	case eightLength:
		return true
	default:
		return false
	}
}

// this function should work as well... but doesnt
// func isOneFourSevenOrEight(outputLetters []string) bool {
// 	switch {
// 	case utils.UnorderedEqual(one, outputLetters):
// 		return true
// 	case utils.UnorderedEqual(four, outputLetters):
// 		return true
// 	case utils.UnorderedEqual(seven, outputLetters):
// 		return true
// 	case utils.UnorderedEqual(eight, outputLetters):
// 		return true
// 	default:
// 		return false
// 	}
// }
