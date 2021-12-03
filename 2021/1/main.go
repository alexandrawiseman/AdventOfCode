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
	numIncreases := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			numIncreases++
		}
	}

	fmt.Println("Part 1: ", numIncreases)
}

func partTwo(nums []int) {
	numIncreases := 0

	for i := 0; i+3 < len(nums); i += 1 {
		groupOne := nums[i] + nums[i+1] + nums[i+2]
		groupTwo := nums[i+1] + nums[i+2] + nums[i+3]
		if groupTwo > groupOne {
			numIncreases++
		}
	}

	fmt.Println("Part 2: ", numIncreases)
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
