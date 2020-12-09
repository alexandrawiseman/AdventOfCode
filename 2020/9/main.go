package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	nums := readFileIntoIntArray()
	partOneResult := partOne(nums)
	partTwo(nums, partOneResult)
}

func partOne(nums []int) int {
	badValue := 0
	preamble := 25

numberIteration:
	for i := preamble; i < len(nums); i++ {
		sumVal := nums[i]
		for x := i - preamble; x < i-1; x++ {
			for y := x + 1; y < i; y++ {
				if nums[x]+nums[y] == sumVal {
					continue numberIteration
				}
			}
		}
		badValue = sumVal
		break
	}

	fmt.Println("Part 1: ", badValue)
	return badValue
}

func partTwo(nums []int, partOneResult int) {
	var smallNum, largeNum int

numberIteration:
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i]
		smallNum = nums[i]
		largeNum = nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]

			if nums[j] < smallNum {
				smallNum = nums[j]
			}

			if nums[j] > largeNum {
				largeNum = nums[j]
			}

			if sum == partOneResult {
				break numberIteration
			}
		}
	}

	fmt.Println("Part 2: ", smallNum+largeNum)
}

func readFileIntoIntArray() []int {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	nums := make([]int, 0, len(lines))

	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}

	return nums
}
