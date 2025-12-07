package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Problem struct {
	nums      []int
	operation string
}

func getProblems(lines []string) []Problem {

	numCols := len(strings.Split(lines[0], " "))

	problems := make([]Problem, numCols)

	for i := 0; i < len(lines); i++ {
		parts := removeEmptyStrings(strings.Split(lines[i], " "))

		for col, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				problems[col].nums = append(problems[col].nums, num)
			} else {
				problems[col].operation = part
			}
		}

	}
	return problems
}

func getProblemsTopBottom(lines []string) []Problem {
	maxLengths := []int{}

	lastLine := lines[len(lines)-1]

	operatorParts := strings.FieldsFunc(lastLine, func(r rune) bool {
		return r == '*' || r == '+'
	})

	maxLineLength := 0
	for _, line := range lines[:len(lines)-1] {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	operations := removeEmptyStrings(strings.Split(lastLine, " "))

	lastMaxLength := maxLineLength - len(strings.TrimSpace(lastLine)) + 1

	for _, part := range operatorParts {
		maxLengths = append(maxLengths, len(part))
	}

	maxLengths[len(maxLengths)-1] = lastMaxLength

	problems := make([]Problem, len(maxLengths))

	startIndex := 0

	for group := 0; group < len(maxLengths); group++ {
		for col := startIndex; col < maxLengths[group]+startIndex; col++ {
			numVal := 0
			for row := 0; row < len(lines)-1; row++ {
				line := lines[row]
				if line[col] != ' ' {
					digit, _ := strconv.Atoi(string(line[col]))
					numVal *= 10
					numVal += digit
				}
			}
			problems[group].nums = append(problems[group].nums, numVal)
		}
		startIndex += maxLengths[group] + 1
		problems[group].operation = operations[group]

	}

	return problems
}

func performOperation(problem Problem) int {
	result := 0
	switch problem.operation {
	case "+":
		for _, num := range problem.nums {
			result += num
		}
	case "*":
		result = 1
		for _, num := range problem.nums {
			result *= num
		}
	}
	return result
}

func Day6() {
	lines, err := openDataFile("day6-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Your code for Day 6 goes here
	problems := getProblems(lines)
	fmt.Println(problems)

	total := 0
	for i, problem := range problems {
		result := performOperation(problem)
		fmt.Printf("Problem %d result: %d\n", i+1, result)
		total += result
	}

	fmt.Println(total)
}

func Day6part2() {
	lines, err := openDataFile("day6-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Your code for Day 6 goes here
	problems := getProblemsTopBottom(lines)
	fmt.Println(problems)

	total := 0
	for i, problem := range problems {
		result := performOperation(problem)
		fmt.Printf("Problem %d result: %d\n", i+1, result)
		total += result
	}

	fmt.Println(total)
}
