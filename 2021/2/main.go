package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Command struct {
	Move  string
	Value int
}

func main() {
	commands := readFileIntoCommandArray()
	partOne(commands)
	partTwo(commands)
}

func partOne(commands []Command) {
	depth := 0
	horizontalPosition := 0

	for _, command := range commands {
		if command.Move == "forward" {
			horizontalPosition += command.Value
		} else if command.Move == "up" {
			depth -= command.Value
		} else {
			depth += command.Value
		}
	}

	fmt.Println("Part 1: ", depth*horizontalPosition)
}

func partTwo(commands []Command) {
	depth := 0
	horizontalPosition := 0
	aim := 0

	for _, command := range commands {
		if command.Move == "forward" {
			horizontalPosition += command.Value
			depth += (aim * command.Value)
		} else if command.Move == "up" {
			aim -= command.Value
		} else {
			aim += command.Value
		}
	}

	fmt.Println("Part 2: ", depth*horizontalPosition)
}

func readFileIntoCommandArray() []Command {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	commands := make([]Command, 0, len(lines))

	for _, line := range lines {
		vals := strings.Split(line, " ")
		n, _ := strconv.Atoi(vals[1])

		c := Command{
			Move:  vals[0],
			Value: n,
		}

		commands = append(commands, c)
	}

	return commands
}
