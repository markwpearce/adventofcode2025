package days

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func getFreshRanges(lines []string) ([][]int, int) {
	freshRanges := [][]int{}
	lastLine := 0
	for i, line := range lines {
		if line == "" {
			lastLine = i
			break
		}
		freshRange := getFreshRange(line)
		freshRanges = append(freshRanges, freshRange)
	}
	return freshRanges, lastLine
}

func getFreshRange(line string) []int {
	parts := strings.Split(strings.TrimSpace(line), "-")

	firstNum, _ := strconv.Atoi(parts[0])
	secondNum, _ := strconv.Atoi(parts[1])
	return []int{firstNum, secondNum}
}

func isNumInRange(num int, freshRange []int) bool {
	return num >= freshRange[0] && num <= freshRange[1]
}

func checkAllRanges(num int, freshRanges [][]int) bool {
	for _, freshRange := range freshRanges {
		if isNumInRange(num, freshRange) {
			return true
		}
	}
	return false
}

func Day5() {
	lines, err := openDataFile("day5-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Example: map with string keys and int values

	freshRanges, lastLine := getFreshRanges(lines)
	freshCount := 0
	for i := lastLine + 1; i < len(lines); i++ {
		num, _ := strconv.Atoi(strings.TrimSpace(lines[i]))

		if checkAllRanges(num, freshRanges) {
			fmt.Printf("Ingredient %d IS fresh\n", num)

			freshCount++
		} else {

			fmt.Printf("Ingredient %d is NOT fresh\n", num)
		}
	}

	fmt.Println(freshCount)

}

func consolidateRanges(freshRanges [][]int) [][]int {
	if len(freshRanges) == 0 {
		return freshRanges
	}

	// First, sort the ranges by their starting values
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	consolidated := [][]int{}
	currentRange := freshRanges[0]

	for i := 1; i < len(freshRanges); i++ {
		nextRange := freshRanges[i]
		if nextRange[0] <= currentRange[1]+1 {
			// Ranges overlap or are contiguous, merge them
			if nextRange[1] > currentRange[1] {
				currentRange[1] = nextRange[1]
			}
		} else {
			// No overlap, add the current range to the list and move to the next
			consolidated = append(consolidated, currentRange)
			currentRange = nextRange
		}
	}
	// Add the last range
	consolidated = append(consolidated, currentRange)

	return consolidated
}

func Day5part2() {
	lines, err := openDataFile("day5-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Example: map with string keys and int values

	freshRanges, _ := getFreshRanges(lines)

	freshRanges = consolidateRanges(freshRanges)

	freshCount := 0

	for _, freshRange := range freshRanges {
		totalInRange := freshRange[1] - freshRange[0] + 1
		fmt.Printf("Range %d-%d has %d ingredients\n", freshRange[0], freshRange[1], totalInRange)
		freshCount += totalInRange
	}

	fmt.Println(freshCount)

}
