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
	partOne(nums)
	partTwo(nums)
}

func partOne(nums []int) {
	resultingNums := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				resultingNums[0] = nums[i]
				resultingNums[1] = nums[j]
			}
		}
	}

	fmt.Println(resultingNums[0] * resultingNums[1])
}

func partTwo(nums []int) {
	resultingNums := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					resultingNums[0] = nums[i]
					resultingNums[1] = nums[j]
					resultingNums[2] = nums[k]
				}
			}
		}
	}

	fmt.Println(resultingNums[0] * resultingNums[1] * resultingNums[2])
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
