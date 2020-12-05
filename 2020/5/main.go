package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	lines := readFileIntoStringArray()
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	largestValue := 0

	for _, line := range lines {
		row, col := convertLineIntoRowAndCol(line)
		value := getUniqueId(row, col)
		if value > largestValue {
			largestValue = value
		}
	}

	fmt.Println(largestValue)
}

func partTwo(lines []string) {
	filledSeats := make([][]bool, 128)
	for i := range filledSeats {
		filledSeats[i] = make([]bool, 8)
	}

	for _, line := range lines {
		row, col := convertLineIntoRowAndCol(line)
		filledSeats[row][col] = true
	}

	row, col := findEmptySeat(filledSeats)

	fmt.Println(getUniqueId(row, col))
}

func convertLineIntoRowAndCol(line string) (int, int) {
	var lowerRowRange float64 = 0
	var upperRowRange float64 = 127
	var lowerColRange float64 = 0
	var upperColRange float64 = 7

	for _, char := range line {
		if char == 'F' {
			upperRowRange = math.Floor((upperRowRange + lowerRowRange) / 2)
		} else if char == 'B' {
			lowerRowRange = math.Ceil((upperRowRange + lowerRowRange) / 2)
		} else if char == 'L' {
			upperColRange = math.Floor((upperColRange + lowerColRange) / 2)
		} else if char == 'R' {
			lowerColRange = math.Ceil((upperColRange + lowerColRange) / 2)
		}
	}

	return int(lowerRowRange), int(lowerColRange)
}

func findEmptySeat(filledSeats [][]bool) (int, int) {
	for row := range filledSeats {
		missingSeats := []int{}
		for col := 0; col < len(filledSeats[row]); col++ {
			val := filledSeats[row][col]
			if !val {
				missingSeats = append(missingSeats, col)
			}
		}

		if len(missingSeats) == 1 {
			return row, missingSeats[0]
		}
	}

	return -1, -1
}

func getUniqueId(row int, col int) int {
	return row*8 + col
}

func readFileIntoStringArray() []string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
