package days

import (
	"bufio"
	"log"
	"os"
)

func openDataFile(fileName string) ([]string, error) {
	file, err := os.Open("./data/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, nil
}
