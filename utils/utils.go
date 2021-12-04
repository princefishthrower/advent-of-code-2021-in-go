package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		log.Fatalf(scanner.Err().Error())
	}
	return lines
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func StringArrayToIntArray(array []string) []int {
	var intArray []int
	for index, item := range array {
		number, err := strconv.ParseInt(item, 10, 0)
		if err != nil {
			log.Fatalf("Could not parse number from string: %s", item)
		}
		intArray[index] = int(number)
	}
	return intArray
}
