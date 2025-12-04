package days

import (
	"fmt"
	"log"
)

func testGroupAround(warehouse []string, row, col int) int {
	rows := len(warehouse)
	cols := len(warehouse[0])
	startRow := maxInt(0, row-1)
	endRow := minInt(rows-1, row+1)
	startCol := maxInt(0, col-1)
	endCol := minInt(cols-1, col+1)

	rollsAround := 0

	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			if r == row && c == col {
				continue
			}
			testChar := warehouse[r][c]
			if testChar == '@' {
				rollsAround++
			}
		}
	}
	return rollsAround
}

func getAccessibleRolls(warehouse []string, maxRolls int) [][]int {
	accessibleRolls := [][]int{}

	for row, line := range warehouse {

		for col, char := range line {
			if char != '@' {
				continue
			}

			rollsAround := testGroupAround(warehouse, row, col)
			fmt.Printf("Testing (%d,%d) rollsAround: %d\n", row, col, rollsAround)

			if rollsAround < maxRolls {
				accessibleRolls = append(accessibleRolls, []int{row, col})
				fmt.Printf("Accessible at (%d,%d)\n", row, col)
			}
		}
	}
	return accessibleRolls
}

func Day4() {
	lines, err := openDataFile("day4-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	accessible := getAccessibleRolls(lines, 4)

	fmt.Println(len(accessible))
}

func Day4part2() {
	lines, err := openDataFile("day4-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	iterations := 0
	totalRolls := 0

	accessible := getAccessibleRolls(lines, 4)
	for len(accessible) > 0 {
		totalRolls += len(accessible)
		for _, coord := range accessible {
			row := coord[0]
			col := coord[1]
			line := lines[row]
			newLine := line[:col] + "x" + line[col+1:]
			lines[row] = newLine
		}
		accessible = getAccessibleRolls(lines, 4)
		iterations++
		fmt.Printf("After iteration %d, accessible: %d\n", iterations, len(accessible))
	}

	fmt.Println(iterations)
	fmt.Println(totalRolls)
}
