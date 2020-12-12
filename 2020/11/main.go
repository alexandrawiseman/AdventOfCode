package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := readFileIntoStringArray()
	valChanged := true
	for valChanged {
		nextRound := make([]string, len(lines))
		copy(nextRound, lines)
		valChanged = false
		for y := 0; y < len(lines); y++ {
			for x := 0; x < len(lines[0]); x++ {
				if lines[y][x] == 'L' && numOfAdjacentOccupiedSeats(lines, x, y) == 0 {
					nextRound[y] = nextRound[y][0:x] + "#" + nextRound[y][x+1:]
					valChanged = true
				} else if lines[y][x] == '#' && numOfAdjacentOccupiedSeats(lines, x, y) >= 4 {
					nextRound[y] = nextRound[y][0:x] + "L" + nextRound[y][x+1:]
					valChanged = true
				}
			}
		}
		copy(lines, nextRound)
	}

	fmt.Println("Part One: ", countOfOccupiedSeats(lines))
}

func partTwo() {
	lines := readFileIntoStringArray()
	valChanged := true
	for valChanged {
		nextRound := make([]string, len(lines))
		copy(nextRound, lines)
		valChanged = false
		for y := 0; y < len(lines); y++ {
			for x := 0; x < len(lines[0]); x++ {
				if lines[y][x] == 'L' && numOccupiedSeatsInAllSurroundingViews(lines, x, y) == 0 {
					nextRound[y] = nextRound[y][0:x] + "#" + nextRound[y][x+1:]
					valChanged = true
				} else if lines[y][x] == '#' && numOccupiedSeatsInAllSurroundingViews(lines, x, y) >= 5 {
					nextRound[y] = nextRound[y][0:x] + "L" + nextRound[y][x+1:]
					valChanged = true
				}
			}
		}
		copy(lines, nextRound)
	}

	fmt.Println("Part Two: ", countOfOccupiedSeats(lines))
}

func isSeat(c byte) bool {
	if c == '#' || c == 'L' {
		return true
	}
	return false
}

func numOccupiedSeatsInAllSurroundingViews(lines []string, x int, y int) int {
	count := 0

	for i := y - 1; i >= 0; i-- {
		if isSeat(lines[i][x]) {
			if lines[i][x] == '#' {
				count++
			}
			break
		}
	}

	for i := y + 1; i < len(lines); i++ {
		if isSeat(lines[i][x]) {
			if lines[i][x] == '#' {
				count++
			}
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		if isSeat(lines[y][i]) {
			if lines[y][i] == '#' {
				count++
			}
			break
		}
	}

	for i := x + 1; i < len(lines[0]); i++ {
		if isSeat(lines[y][i]) {
			if lines[y][i] == '#' {
				count++
			}
			break
		}
	}

	xVal := x + 1
	yVal := y + 1
	for xVal < len(lines[0]) && yVal < len(lines) {
		if isSeat(lines[yVal][xVal]) {
			if lines[yVal][xVal] == '#' {
				count++
			}
			break
		}
		xVal++
		yVal++
	}

	xVal = x - 1
	yVal = y - 1
	for xVal >= 0 && yVal >= 0 {
		if isSeat(lines[yVal][xVal]) {
			if lines[yVal][xVal] == '#' {
				count++
			}
			break
		}
		xVal--
		yVal--
	}

	xVal = x - 1
	yVal = y + 1
	for xVal >= 0 && yVal < len(lines) {
		if isSeat(lines[yVal][xVal]) {
			if lines[yVal][xVal] == '#' {
				count++
			}
			break
		}
		xVal--
		yVal++
	}

	xVal = x + 1
	yVal = y - 1
	for xVal < len(lines[0]) && yVal >= 0 {
		if isSeat(lines[yVal][xVal]) {
			if lines[yVal][xVal] == '#' {
				count++
			}
			break
		}
		xVal++
		yVal--
	}

	return count
}

func numOfAdjacentOccupiedSeats(lines []string, x int, y int) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(i == x && j == y) && (i >= 0 && i < len(lines[0]) && j >= 0 && j < len(lines)) {
				if lines[j][i] == '#' {
					count++
				}
			}
		}
	}
	return count
}

func countOfOccupiedSeats(lines []string) int {
	count := 0
	for _, line := range lines {
		for _, c := range line {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func readFileIntoStringArray() []string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
