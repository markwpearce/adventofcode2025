package days

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func splitStringToNumbers(s string) []int {
	var numbers []int
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(string(s[i]))
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func getMaxJolts(bank string) int {
	nums := splitStringToNumbers(bank)
	first := nums[len(nums)-2]
	second := nums[len(nums)-1]
	currentVal := getJolts([]int{first, second})

	for i := len(nums) - 3; i >= 0; i-- {
		joltsWithFirst := getJolts([]int{nums[i], first})
		joltsWithSecond := getJolts([]int{nums[i], second})
		startingSecond := second

		if joltsWithFirst > currentVal {
			second = first
			first = nums[i]
			currentVal = getJolts([]int{first, second})

		}
		if joltsWithSecond > currentVal {
			first = nums[i]
			second = startingSecond
			currentVal = getJolts([]int{first, second})

		}
	}
	return currentVal
}

func getMaxJolts2(bank string, digits int) int {
	nums := splitStringToNumbers(bank)

	currentChoice := make([]int, digits)
	for i := 0; i < digits; i++ {
		currentChoice[i] = nums[len(nums)-digits+i]
	}

	currentVal := getJolts(currentChoice)
	result := currentChoice

	for i := len(nums) - digits - 1; i >= 0; i-- {

		for j := 0; j < digits; j++ {
			testNums := putInFirstPlaceRemoveAtIndex(nums[i], currentChoice, j)
			joltsWithReplace := getJolts(testNums)
			if joltsWithReplace > currentVal {
				result = testNums
				currentVal = joltsWithReplace
			}
		}
		currentChoice = result

	}

	return currentVal
}

func getJolts(nums []int) int {
	value := 0
	for i := len(nums) - 1; i >= 0; i-- {
		value += nums[i] * int(math.Pow10(len(nums)-1-i))
	}
	return value
}

func putInFirstPlaceRemoveAtIndex(newNum int, nums []int, index int) []int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)

	for i := index; i >= 1; i-- {
		newNums[i] = newNums[i-1]
	}
	newNums[0] = newNum
	return newNums
}

func Day3() {
	lines, err := openDataFile("day3-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	sum := 0
	for _, line := range lines {
		result := getMaxJolts2(line, 12)
		sum += result
		fmt.Println("Result:", result)
	}
	fmt.Println(sum)
}
