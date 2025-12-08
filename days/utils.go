package days

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
