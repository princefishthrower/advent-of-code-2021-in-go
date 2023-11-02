package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.ReadLines("input.txt")

	var heightMap [][]int
	// for each note, get the segment output part of the note
	// then count the 1s,4s,7s, and 8s in the outputs
	maxY := 0
	for _, line := range lines {
		heightLineStrings := strings.Split(line, "")
		heightLine := utils.StringArrayToIntArray(heightLineStrings)
		heightMap = append(heightMap, heightLine)
		maxY++
	}

	totalRisk := 0
	for x := 0; x < len(heightLine); x++ {
		for y := 0; y < yHeight; y++ {
			risk := getRiskIfMin(heightmap, maxY, x, y)
			totalRisk += risk
		}
	}

	// solution is the totalRisk (sum of risks at each min)
	fmt.Println(totalRisk)
}

func getRiskIfMin(heightMap [][]int, maxX int, maxY int, x int, y int) int {

	valueAtPoint := heightMap[x][y]

	switch {
	// edgecase 1 - top left corner
	case x == 0 && y == 0 && valueAtPoint < heightMap[x+1][y] &&
		valueAtPoint < heightMap[x][y-1]:
		return valueAtPoint + 1

	// edgecase 2 - top right corner
	case x == maxX && y == 0 && valueAtPoint < heightMap[x-1][y] &&
		valueAtPoint < heightMap[x][y-1]:
		return valueAtPoint + 1

	// edgecase 3 - bottom left corner
	case x == 0 && y == maxY && valueAtPoint < heightMap[x+1][y] &&
		valueAtPoint < heightMap[x][y+1]:
		return valueAtPoint + 1

	// edgecase 4 - bottom right corner
	case x == maxX && y == maxY && valueAtPoint < heightMap[x-1][y] &&
		valueAtPoint < heightMap[x][y+1]:
		return valueAtPoint + 1

	// edgease 5 - top row
	case x == 0 && y != 0 && y != maxY-1 && valueAtPoint < heightMap[x-1][y] && valueAtPoint < heightMap[x+1][y] && valueAtPoint < heightMap[x][y-1]:
		return valueAtPoint + 1

	// normal evaluation
	case valueAtPoint < heightMap[x][y+1] &&
		valueAtPoint < heightMap[x][y-1] &&
		valueAtPoint < heightMap[x-1][y] &&
		valueAtPoint < heightMap[x+1][y]:
		return valueAtPoint + 1

	default:
		return 0
	}
}
