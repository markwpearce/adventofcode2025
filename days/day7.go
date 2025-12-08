package days

import (
	"fmt"
	"log"
	"slices"
)

type TachyonTree struct {
	startCol        int
	splitterIndexes [][]int
}

func parseTachyonTree(lines []string) TachyonTree {
	cols := len(lines[0])

	tree := TachyonTree{}

	for i := 0; i < cols; i++ {
		if lines[0][i] == 'S' {
			tree.startCol = i
			break
		}
	}

	for row := 1; row < len(lines)-1; row++ {
		rowSplits := []int{}
		line := lines[row]
		for col := 0; col < cols; col++ {
			if line[col] == '^' {
				rowSplits = append(rowSplits, col)
			}
		}
		tree.splitterIndexes = append(tree.splitterIndexes, rowSplits)
	}

	return tree
}

func flowBeam(tree TachyonTree) int {

	curBeamPositions := []int{tree.startCol}
	hits := [][]int{}
	for row, splitters := range tree.splitterIndexes {
		if row%2 == 0 {
			continue
		}
		nextBeamPositions := []int{}
		for _, pos := range curBeamPositions {
			hit := slices.Contains(splitters, pos)

			if hit {
				// Handle splitter hit - add split beam positions
				hits = append(hits, []int{row, pos})
				fmt.Printf("Beam hit splitter at row %d, col %d\n", row, pos)
				nextBeamPositions = append(nextBeamPositions, pos-1, pos+1)

			} else {
				// Beam continues straight
				nextBeamPositions = append(nextBeamPositions, pos)
			}
		}

		// Remove duplicates from nextBeamPositions

		nextBeamPositions = makeUnique(nextBeamPositions)
		curBeamPositions = nextBeamPositions
	}
	return len(hits)
}

func flowBeam2(tree TachyonTree) (int, int) {
	beamsByIndex := make(map[int]int)
	beamsByIndex[tree.startCol] = 1
	totalTimelines := 1

	for row, splitters := range tree.splitterIndexes {
		if row%2 == 0 {
			continue
		}
		hitsAtRow := make(map[int]int)
		for pos, beamsAtPos := range beamsByIndex {

			hit := beamsAtPos > 0 && slices.Contains(splitters, pos)

			if hit {
				hitsAtRow[pos] = beamsAtPos
				totalTimelines += beamsAtPos
			}
		}
		fmt.Println("Row", row+2, " hits at row: ", hitsAtRow)

		for pos, hitcountAtPos := range hitsAtRow {
			beamsAtPos := beamsByIndex[pos]
			leftOptionBeams := beamsByIndex[pos-1]
			beamsByIndex[pos-1] = leftOptionBeams + hitcountAtPos
			rightOptionBeams := beamsByIndex[pos+1]
			beamsByIndex[pos+1] = rightOptionBeams + hitcountAtPos
			beamsByIndex[pos] = beamsAtPos - hitcountAtPos
		}

		// Remove duplicates from nextBeamPositions
		fmt.Println("at row", row+2, " number of current timelines: ", totalTimelines)
	}
	sum := 0
	for _, count := range beamsByIndex {
		sum += count
	}

	return sum, totalTimelines
}

func Day7() {
	lines, err := openDataFile("day7-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Your code for Day 7 goes here
	tree := parseTachyonTree(lines)
	fmt.Println(tree)
	hits := flowBeam(tree)
	fmt.Println("Number of hits:", hits)

}

func Day7part2() {
	lines, err := openDataFile("day7-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Your code for Day 7 goes here
	tree := parseTachyonTree(lines)
	fmt.Println(tree)
	hits, timelines := flowBeam2(tree)
	fmt.Println("Number of hits:", hits)
	fmt.Println("Total beams:", timelines)

}
