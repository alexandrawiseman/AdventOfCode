package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := readFileIntoIntArray()
	partOne(nums)
	partTwo(nums)
}

func partOne(nums []int) {
	// Wall is 0 jolts
	nums = append(nums, 0)

	sort.Ints(nums)

	// Built in adapter is three highest than the highest adapter
	nums = append(nums, nums[len(nums)-1]+3)

	countOfOnes := 0
	countOfThrees := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			countOfOnes++
		} else if nums[i+1]-nums[i] == 3 {
			countOfThrees++
		} else {
			log.Fatalf("Bad puzzle input")
		}
	}

	fmt.Println("Part One Result: ", countOfOnes*countOfThrees)
}

func partTwo(nums []int) {
	// Wall is 0 jolts
	nums = append(nums, 0)

	sort.Ints(nums)

	// Built in adapter is three highest than the highest adapter
	nums = append(nums, nums[len(nums)-1]+3)

	seen := make(map[int]int)

	result := getVariants(nums, 0, seen)

	fmt.Println("Part Two: ", result)
}

func getVariants(nums []int, index int, seen map[int]int) int {
	if val, ok := seen[index]; ok {
		return val
	}

	if index == len(nums)-1 {
		return 1
	}

	count := 0
	for i := index + 1; i < len(nums); i++ {
		if nums[i]-nums[index] <= 3 {
			count += getVariants(nums, i, seen)
		}
	}

	seen[index] = count

	return count
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
