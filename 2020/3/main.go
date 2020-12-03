package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	values := readFileIntoStringArray()
	partOne(values)
	partTwo(values)
}

type Slope struct {
	X int
	Y int
}

func partOne(lines []string) {
	var xCoord, yCoord, countOfTrees int

	for yCoord < len(lines) {
		line := lines[yCoord]
		locationValue := line[xCoord]

		if locationValue == '#' {
			countOfTrees++
		}

		yCoord++
		xCoord += 3

		if xCoord >= len(lines[0]) {
			xCoord = xCoord - len(lines[0])
		}
	}

	fmt.Println(countOfTrees)
}

func partTwo(lines []string) {
	slopes := []Slope{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 3,
			Y: 1,
		},
		{
			X: 5,
			Y: 1,
		},
		{
			X: 7,
			Y: 1,
		},
		{
			X: 1,
			Y: 2,
		},
	}

	total := 1

	for _, slope := range slopes {
		var xCoord, yCoord, countOfTrees int

		for yCoord < len(lines) {
			line := lines[yCoord]
			locationValue := line[xCoord]

			if locationValue == '#' {
				countOfTrees++
			}

			yCoord += slope.Y
			xCoord += slope.X

			if xCoord >= len(lines[0]) {
				xCoord = xCoord - len(lines[0])
			}
		}

		total = total * countOfTrees
	}

	fmt.Println(total)

}

func readFileIntoStringArray() []string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
