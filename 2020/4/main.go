package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	index := 0
	b := 0
	countOfValidPassports := 0

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for scanner.Scan() {
		if scanner.Text() == "" {
			if b == 127 {
				countOfValidPassports++
			}

			b = 0
		}

		for i := range fields {
			if strings.Contains(scanner.Text(), fields[i]) {
				b |= (1 << (i))
			}
		}

		index++
	}

	if b == 127 {
		countOfValidPassports++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countOfValidPassports)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index := 0
	b := 0
	countOfValidPassports := 0

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for scanner.Scan() {
		if scanner.Text() == "" {
			if b == 127 {
				countOfValidPassports++
			}

			b = 0
		}

		for i := range fields {
			if strings.Contains(scanner.Text(), fields[i]) {
				var value string
				substrings := strings.Split(scanner.Text(), " ")
				for _, str := range substrings {
					if strings.Contains(str, fields[i]) {
						value = strings.Split(str, ":")[1]
						break
					}
				}

				if isValid(value, fields[i]) {
					b |= (1 << (i))
				}
			}
		}

		index++
	}

	if b == 127 {
		countOfValidPassports++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countOfValidPassports)
}

func isValid(str string, field string) bool {
	switch field {
	case "byr":
		val, _ := strconv.Atoi(str)
		if len(str) == 4 && val >= 1920 && val <= 2002 {
			return true
		}
	case "iyr":
		val, _ := strconv.Atoi(str)
		if len(str) == 4 && val >= 2010 && val <= 2020 {
			return true
		}
	case "eyr":
		val, _ := strconv.Atoi(str)
		if len(str) == 4 && val >= 2020 && val <= 2030 {
			return true
		}
	case "hgt":
		matchCm, _ := regexp.MatchString("[0-9]+cm", str)
		if matchCm {
			val, _ := strconv.Atoi(str[0:(len(str) - 2)])
			if val >= 150 && val <= 193 {
				return true
			}
			break
		}

		matchIn, _ := regexp.MatchString("[0-9]+in", str)
		if matchIn {
			val, _ := strconv.Atoi(str[0:(len(str) - 2)])
			if val >= 59 && val <= 76 {
				return true
			}
		}
	case "hcl":
		match, _ := regexp.MatchString("[#]([0-9]|[a-f]){6}", str)
		if match {
			return true
		}
	case "ecl":
		acceptableVals := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, acceptableVal := range acceptableVals {
			if str == acceptableVal {
				return true
			}
		}
	case "pid":
		match, _ := regexp.MatchString("^([0-9]){9}$", str)
		if match {
			return true
		}
	}

	return false
}
