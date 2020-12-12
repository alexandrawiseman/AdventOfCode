package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	Action rune
	Amount int
}

func main() {
	instructions := readFileIntoInstructionArray()
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []Instruction) {
	direction := 'E'
	var locX, locY int

	for _, instruction := range instructions {
		if instruction.Action == 'F' {
			instruction.Action = direction
		}

		direction, locX, locY = getupdateDDegreesDirectionAndLocationPartOne(instruction, direction, locX, locY)
	}

	manhattanDistance := math.Abs(float64(locX)) + math.Abs(float64(locY))
	fmt.Println("Part One: ", manhattanDistance)
}

func partTwo(instructions []Instruction) {
	var locX, locY int
	wayPointX := 10
	wayPointY := 1

	for _, instruction := range instructions {
		switch instruction.Action {
		case 'L':
			x := wayPointX
			y := wayPointY

			switch instruction.Amount % 360 {
			case 90:
				wayPointX = -y
				wayPointY = x
			case 180:
				wayPointX = -x
				wayPointY = -y
			case 270:
				wayPointX = y
				wayPointY = -x
			}
		case 'R':
			x := wayPointX
			y := wayPointY

			switch instruction.Amount % 360 {
			case 90:
				wayPointX = y
				wayPointY = -x
			case 180:
				wayPointX = -x
				wayPointY = -y
			case 270:
				wayPointX = -y
				wayPointY = x
			}
		case 'N':
			wayPointY += instruction.Amount
		case 'S':
			wayPointY -= instruction.Amount
		case 'W':
			wayPointX -= instruction.Amount
		case 'E':
			wayPointX += instruction.Amount
		case 'F':
			locX += instruction.Amount * wayPointX
			locY += instruction.Amount * wayPointY
		}
	}

	manhattanDistance := math.Abs(float64(locX)) + math.Abs(float64(locY))
	fmt.Println("Part Two: ", manhattanDistance)
}

func getupdateDDegreesDirectionAndLocationPartOne(instruction Instruction, direction rune, locX int, locY int) (rune, int, int) {
	degreesToDirection := map[int]rune{
		0:    'E',
		90:   'S',
		180:  'W',
		270:  'N',
		-90:  'N',
		-180: 'W',
		-270: 'S',
	}

	directionToDegrees := map[rune]int{
		'E': 0,
		'S': 90,
		'W': 180,
		'N': 270,
	}

	degrees := directionToDegrees[direction]

	switch instruction.Action {
	case 'L':
		degrees = (degrees - instruction.Amount) % 360
		direction = degreesToDirection[degrees]
	case 'R':
		degrees = (degrees + instruction.Amount) % 360
		direction = degreesToDirection[degrees]
	case 'N':
		locY += instruction.Amount
	case 'S':
		locY -= instruction.Amount
	case 'W':
		locX -= instruction.Amount
	case 'E':
		locX += instruction.Amount
	}

	return direction, locX, locY
}

func readFileIntoInstructionArray() []Instruction {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	instructions := []Instruction{}
	for _, line := range lines {
		amount, _ := strconv.Atoi(line[1:])

		instructions = append(instructions, Instruction{
			Action: rune(line[0]),
			Amount: amount,
		})
	}

	return instructions
}
