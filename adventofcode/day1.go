package adventofcode

import (
	"fmt"
	"strings"
)

func Day1(input string) (int, int) {
	var floor int = 0
	var firstBasement int = -1

	for index, char := range strings.Split(input, "") {
		if char == "(" {
			floor++
		} else if char == ")" {
			floor--
		} else {
			fmt.Print(fmt.Sprintf("ignoring invalid character %s", char))
		}

		if firstBasement == -1 && floor == -1 {
			firstBasement = index + 1
		}
	}

	return floor, firstBasement
}
