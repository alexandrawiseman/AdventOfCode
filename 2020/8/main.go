package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	Command string
	Value   int
}

func main() {
	instructions := readFileIntoInstructionArray()
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []Instruction) {
	seenInstructions := make([]bool, len(instructions))
	accumulator := 0
	i := 0

	for true {
		if seenInstructions[i] {
			break
		}

		seenInstructions[i] = true
		if instructions[i].Command == "jmp" {
			i += instructions[i].Value
		} else {
			if instructions[i].Command == "acc" {
				accumulator += instructions[i].Value
			}
			i++
		}
	}

	fmt.Println(accumulator)
}

func partTwo(instructions []Instruction) {
	accumulator := 0
	for changedIndex := 0; changedIndex < len(instructions); changedIndex++ {
		if instructions[changedIndex].Command == "acc" {
			continue
		} else if instructions[changedIndex].Command == "nop" {
			instructions[changedIndex].Command = "jmp"
		} else {
			instructions[changedIndex].Command = "nop"
		}

		i := 0
		seenInstructions := make([]bool, len(instructions))
		accumulator = 0

		for i < len(instructions) {
			if seenInstructions[i] {
				break
			}

			seenInstructions[i] = true
			if instructions[i].Command == "jmp" {
				i += instructions[i].Value
			} else {
				if instructions[i].Command == "acc" {
					accumulator += instructions[i].Value
				}
				i++
			}
		}

		if i == len(instructions) {
			break
		}

		// Undo change to instructions for next iteration
		if instructions[changedIndex].Command == "jmp" {
			instructions[changedIndex].Command = "nop"
		} else {
			instructions[changedIndex].Command = "jmp"
		}
	}

	fmt.Println(accumulator)
}

func readFileIntoInstructionArray() []Instruction {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	instructions := []Instruction{}
	for _, line := range lines {
		values := strings.Split(line, " ")
		var instruction Instruction
		instruction.Command = values[0]
		instruction.Value, _ = strconv.Atoi(values[1])
		instructions = append(instructions, instruction)
	}

	return instructions
}
