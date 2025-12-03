package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isInvalid1(id int) bool {
	strNum := strconv.Itoa(id)

	if len(strNum)%2 == 0 {
		firstHalf := strNum[:len(strNum)/2]
		secondHalf := strNum[len(strNum)/2:]
		if firstHalf == secondHalf {
			return true
		}
	}
	return false
}

func splitIntoEqualParts(s string) [][]string {
	var result [][]string
	length := len(s)

	for partSize := 1; partSize <= length; partSize++ {
		if length%partSize != 0 {
			continue
		}

		var parts []string
		for i := 0; i < length; i += partSize {
			parts = append(parts, s[i:i+partSize])
		}
		result = append(result, parts)
	}

	return result
}

func isInvalid2(id int) bool {
	strNum := strconv.Itoa(id)

	equalParts := splitIntoEqualParts(strNum)

	for _, parts := range equalParts {
		allEqual := true
		for i := 1; i < len(parts); i++ {
			if parts[i] != parts[0] {
				allEqual = false
				break
			}
		}
		if allEqual && len(parts) > 1 {
			return true
		}
	}
	return false
}

func Day2() {
	lines, err := openDataFile("day2-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	ranges := []string{}

	for _, line := range lines {
		ranges = append(ranges, strings.Split(line, ",")...)
	}

	invalidIds := []int{}

	for _, idRange := range ranges {
		nums := strings.Split(idRange, "-")
		if len(nums) == 2 {
			first, err1 := strconv.Atoi(strings.TrimSpace((nums[0])))
			second, err2 := strconv.Atoi(strings.TrimSpace((nums[1])))
			if err1 == nil && err2 == nil {
				// Use first and second as needed
				fmt.Println(first, second)
			}

			for i := first; i <= second; i++ {
				if isInvalid2(i) {
					invalidIds = append(invalidIds, i)
				}
			}

			fmt.Println("For Range::", idRange, ", Invalid IDs:", invalidIds)
		}
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
	}
	fmt.Println("Sum of invalid IDs:", sum)

}
