package days

import "math"

func removeEmptyStrings(parts []string) []string {
	var filtered []string
	for _, s := range parts {
		if s != "" {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func makeUnique(input []int) []int {
	seen := make(map[int]bool)
	uniquePositions := []int{}
	for _, pos := range input {
		if !seen[pos] {
			seen[pos] = true
			uniquePositions = append(uniquePositions, pos)
		}
	}
	return uniquePositions
}

func distance(a []int, b []int) float64 {
	if len(a) != len(b) {
		return -1
	}
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	dz := a[2] - b[2]
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
