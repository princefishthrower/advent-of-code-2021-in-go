package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

type VentLine struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func main() {
	ventLines := parseInput()

	// for every point on radar, calculate if there is at least one intersect
	count := countIntersections(ventLines)

	// solution is count of the intersects - print the solution!
	fmt.Println(count)
}

func parseInput() []VentLine {
	lines := utils.ReadLines("input.txt")

	var ventLines []VentLine
	// parse input
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		coordsStartString := strings.Split(points[0], ",")
		coordsStart := utils.StringArrayToIntArray(coordsStartString)
		coordsEndString := strings.Split(points[1], ",")
		coordsEnd := utils.StringArrayToIntArray(coordsEndString)

		// for first part of challenge, only consider horizontal or vertical lines:
		// (x1 == x2) or (y1 == y2)
		if coordsStart[0] == coordsEnd[0] || coordsStart[1] == coordsEnd[1] {
			ventLine := VentLine{coordsStart[0], coordsStart[1], coordsEnd[0], coordsEnd[1]}
			ventLines = append(ventLines, ventLine)
		}
	}

	return ventLines
}

// func getMaxCoords(ventLines []VentLine) (int, int) {
// 	var allX []int
// 	var allY []int
// 	for _, ventLine := range ventLines {
// 		allX = append(allX, ventLine.X1)
// 		allX = append(allX, ventLine.X2)
// 		allY = append(allY, ventLine.Y1)
// 		allY = append(allY, ventLine.Y2)
// 	}
// 	return utils.Max(allX), utils.Max(allY)
// }

// while the challenge refers to 'overlap', its really the intersect
func countIntersections(ventLines []VentLine) int {
	count := 0
	for firstIndex, firstLine := range ventLines {
		for secondIndex, secondLine := range ventLines {
			// if we're looking at the same VentLine, skip
			// we need to ensure they are different VentLines!
			if firstIndex == secondIndex {
				continue
			}
			// intersect calculation
			if doLinesIntersect(firstLine, secondLine) {
				count++
			}
		}
	}
	return count
}

// jeez, cramer's rule to solve a friggen advent of code...
func doLinesIntersect(firstVentLine VentLine, secondVentLines VentLine) bool {
	a := firstVentLine.Y1 - firstVentLine.Y2     //y1 - y2
	b := firstVentLine.X1 - firstVentLine.X2     //x1 - x2
	c := secondVentLines.Y1 - secondVentLines.Y2 //y3 - y4
	d := secondVentLines.X1 - secondVentLines.X2 //x3 - x4
	// e := a*firstVentLine.X1 - b*firstVentLine.Y1   //(y1 - y2)x1 - (x1 - x2)y1
	// f := c*secondVentLine.X2 - d*secondVentLine.Y2 //(y3 - y4)x3 - (x3 - x4)y3

	determinant := a*d - b*c

	return determinant != 0
	//If determinant = 0, there is no solution to the equation,
	//this shows the  lines are parallel

	//There is a solution to the equation
	//this shows the lines intersect.
	// result[0] = (b * f - e * d) / determinant;
	// result[1] = (a * f - e * c) / determinant;

}
