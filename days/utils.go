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
