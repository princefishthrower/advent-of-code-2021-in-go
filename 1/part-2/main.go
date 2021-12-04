package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var text = scanner.Text()
		number, err := strconv.ParseInt(text, 10, 0)
		if err != nil {
			log.Fatalf("Could not parse number from string: %s", text)
		}
		lines = append(lines, int(number))
	}
	return lines, scanner.Err()
}

func main() {
	var count = 0
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	// sliding window of 3 over the lines
	var prevSum = utils.Sum(lines[0:3])
	for i := 1; i < len(lines); i++ {
		var currSum = utils.Sum(lines[i : i+3])
		if currSum > prevSum {
			count++
		}
		prevSum = currSum
	}

	fmt.Println(count)
}
