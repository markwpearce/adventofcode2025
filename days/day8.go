package days

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func getPoints(lines []string) [][]int {

	points := [][]int{}

	for row, line := range lines {
		parts := strings.Split(line, ",")

		if len(parts) != 3 {
			log.Fatalf("invalid point format at row %d: %s", row, line)
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, []int{x, y, z})
	}
	return points
}

func findShortestDistance(point []int, points [][]int) (int, float64) {
	minDistance := -1.0
	clostestPointIndex := -1
	for i, otherPoint := range points {
		dist := distance(point, otherPoint)

		if (minDistance < 0 || dist < minDistance) && dist > 0 {
			minDistance = dist
			clostestPointIndex = i
		}
	}
	return clostestPointIndex, minDistance
}

type Link struct {
	distance float64
	pointA   int
	pointB   int
}

func getChains(points [][]int, maxLinks int) ([][]int, []int) {
	chains := [][]int{}

	pointCount := len(points)

	closestIndexes := []Link{}

	for i, pointA := range points {
		for j, pointB := range points {
			if i == j {
				continue
			}
			dist := distance(pointA, pointB)
			closestIndexes = append(closestIndexes, Link{distance: dist, pointA: i, pointB: j})
		}
	}

	sort.Slice(closestIndexes, func(i, j int) bool {
		return closestIndexes[i].distance < closestIndexes[j].distance
	})

	filteredLinks := []Link{}
	for i, link := range closestIndexes {
		if i > 0 {
			prev := closestIndexes[i-1]
			if link.pointA == prev.pointB && link.pointB == prev.pointA {
				continue
			}
		}
		filteredLinks = append(filteredLinks, link)
	}
	closestIndexes = filteredLinks

	fmt.Println("Closest links:")

	for _, link := range closestIndexes {
		fmt.Printf("Point %d -> Point %d (distance: %.2f)\n", link.pointA, link.pointB, link.distance)
	}

	if maxLinks < 0 {
		maxLinks = len(closestIndexes)
	}

	lastConnection := []int{}

	for i := 0; i < min(maxLinks, len(closestIndexes)); i++ {
		link := closestIndexes[i]
		fmt.Printf("Processing link: Point %d -> Point %d (distance: %.2f)\n", link.pointA, link.pointB, link.distance)
		chainIndexWithA := -1
		chainIndexWithB := -1
		for chainIndex, chain := range chains {
			if slices.Contains(chain, link.pointA) {
				chainIndexWithA = chainIndex

			}

			if slices.Contains(chain, link.pointB) {
				chainIndexWithB = chainIndex
			}
		}

		if chainIndexWithA >= 0 && chainIndexWithB >= 0 {
			if chainIndexWithA != chainIndexWithB {
				// Merge chains
				chains[chainIndexWithA] = append(chains[chainIndexWithA], chains[chainIndexWithB]...)
				// Remove the merged chain
				chains = append(chains[:chainIndexWithB], chains[chainIndexWithB+1:]...)
			}
		} else if chainIndexWithA >= 0 {
			chains[chainIndexWithA] = append(chains[chainIndexWithA], link.pointB)
		} else if chainIndexWithB >= 0 {
			chains[chainIndexWithB] = append(chains[chainIndexWithB], link.pointA)
		} else {
			// Neither point is in any chain, create a new chain
			chains = append(chains, []int{link.pointA, link.pointB})
		}
		if len(chains[0]) == pointCount {
			fmt.Println("All points connected in a single chain.")
			lastConnection = []int{link.pointA, link.pointB}
			break

		}

	}
	/*for i := maxLinks; i < len(closestIndexes); i++ {
		link := closestIndexes[i]
		found := false
		for _, chain := range chains {
			if slices.Contains(chain, link.pointA) {
				found = true
				break
			}
		}
		if !found {
			chains = append(chains, []int{link.pointA})
		}
	}*/
	return chains, lastConnection
}

func Day8() {
	lines, err := openDataFile("day8-1.txt")
	if err != nil {
		log.Fatalf("could not read input: %v", err)
	}

	fmt.Println(lines)

	points := getPoints(lines)
	fmt.Printf("Parsed points: %v\n", points)

	chains, lastConnection := getChains(points, -1)

	sort.Slice(chains, func(i, j int) bool {
		return len(chains[i]) > len(chains[j])
	})

	fmt.Printf("Sorted Chains: %v\n", chains)

	product := 1
	for i, chain := range chains {
		chainLength := len(chain)
		fmt.Printf("Chain length: %d\n", chainLength)
		product *= chainLength
		if i >= 2 {
			break
		}
	}

	fmt.Printf("Product of top 3 chain lengths: %d\n", product)

	if len(chains) == 1 {
		lastPoint := lastConnection[1]
		secondLastPoint := lastConnection[0]

		xProduct := points[lastPoint][0] * points[secondLastPoint][0]

		fmt.Println(lastConnection, xProduct)
	}

}
