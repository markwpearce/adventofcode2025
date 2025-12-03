package days

import (
	"fmt"
	"log"
)

func Day1() {
	lines, err := openDataFile("day1-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	start := 50
	total := 100

	current := start

	zeroes := 0
	for _, element := range lines {
		direction := element[0:1]
		number := 0
		fmt.Sscanf(element[1:], "%d", &number)
		crossings := 0
		if direction == "L" {
			number = -number
		}

		next := current + number

		if next >= total {
			for next > total {
				next -= total
				crossings++
			}
			if next == total {
				next = 0
			}

		} else if next < 0 {
			if current == 0 {
				crossings--
			}
			for next < 0 {
				next += total
				crossings++
			}

			//next = (next % total) + total

		}

		if next == 0 {
			crossings++
		}
		if crossings < 0 {
			crossings = 0
		}

		current = next
		zeroes += crossings

		fmt.Println("After ", element, " - Current position:", current, " - Crossings:", crossings)
	}

	fmt.Println("total zeroes", zeroes)
}
