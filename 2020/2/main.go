package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := readFileIntoStringArray()

	validPasswords := 0
	for _, line := range lines {
		lowerRange, upperRange, letter, password := retrieveValuesFromLine(line)

		count := 0
		for i := 0; i < len(password); i++ {
			if string(password[i]) == letter {
				count++
			}
		}

		if count >= lowerRange && count <= upperRange {
			validPasswords++
		}

	}

	fmt.Println("Valid passwords: ", validPasswords)
}

func partTwo() {
	lines := readFileIntoStringArray()

	validPasswords := 0
	for _, line := range lines {
		lowerRange, upperRange, letter, password := retrieveValuesFromLine(line)

		if exclusivelyOneStringEqualsValue(string(password[lowerRange-1]), string(password[upperRange-1]), letter) {
			validPasswords++
		}
	}

	fmt.Println("Valid passwords: ", validPasswords)
}

func readFileIntoStringArray() []string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

func retrieveValuesFromLine(str string) (lowerRange int, upperRange int, letter string, password string) {
	vals := strings.Split(str, "-")
	lowerRange, _ = strconv.Atoi(vals[0])

	vals = strings.Split(vals[1], " ")
	upperRange, _ = strconv.Atoi(vals[0])
	letter = strings.Trim(vals[1], ":")
	password = vals[2]

	return
}

func exclusivelyOneStringEqualsValue(str1 string, str2 string, value string) bool {
	if (str1 == value && str2 != value) || (str1 != value && str2 == value) {
		return true
	}
	return false
}
