package main

import (
	"advent/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// this is nearly the same as part 1, but we swap out the depth rules for this new "aim"

func main() {

	// read input as string
	stringLines := utils.ReadLines("input.txt")

	// start tracking depth and horizontal position
	depth := 0
	position := 0
	aim := 0

	for _, stringLine := range stringLines {
		// split by space to get command and amount
		commandEntries := strings.Split(stringLine, " ")
		if len(commandEntries) != 2 {
			log.Fatalf("Invalid command entries: %s", commandEntries)
		}

		command := commandEntries[0]
		amount := commandEntries[1]

		amountInt, err := strconv.ParseInt(amount, 10, 0)
		if err != nil {
			log.Fatalf("Cannot convert to int: %s", amount)
		}

		amountNumber := int(amountInt)

		switch command {
		case "forward":
			position += amountNumber
			// same code as part 1, except this rule:
			depth += (aim * amountNumber)
		case "down":
			aim += amountNumber
		case "up":
			aim -= amountNumber
		default:
			log.Fatalf("Invalid command string: %s", command)
		}
	}

	// solution is the depth multiplied by the position
	solution := depth * position

	// print the solution!
	fmt.Println(solution)
}
