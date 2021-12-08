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
	for _, item := range array {
		number, err := strconv.ParseInt(item, 10, 0)
		if err != nil {
			log.Fatalf("Could not parse number from string: %s", item)
		}
		intArray = append(intArray, int(number))
	}
	return intArray
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func Min(array []int) int {
	var min int = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
