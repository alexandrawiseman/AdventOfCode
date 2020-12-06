package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	seenLetters := make([]bool, 26)

	for scanner.Scan() {
		if scanner.Text() == "" {
			seenLetters = make([]bool, 26)
		}

		for i := range scanner.Text() {
			index := scanner.Text()[i] - 'a'
			hasLetter := seenLetters[index]
			if !hasLetter {
				count++
				seenLetters[index] = true
			}
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	var seenLetters *int

	for scanner.Scan() {
		if scanner.Text() == "" {
			count += countNumberOfOnes(*seenLetters)
			seenLetters = nil
			continue
		}

		personsAnswers := 0b0
		for i := range scanner.Text() {
			index := scanner.Text()[i] - 'a'
			personsAnswers = personsAnswers | (1 << index)
		}

		if seenLetters == nil {
			seenLetters = &personsAnswers
		} else {
			*seenLetters &= personsAnswers
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count += countNumberOfOnes(*seenLetters)

	fmt.Println(count)
}

func countNumberOfOnes(input int) int {
	count := 0
	str := fmt.Sprintf("%026b", input)
	for _, c := range str {
		if c == '1' {
			count++
		}
	}

	return count
}
