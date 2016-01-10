package adventofcode

import (
	"errors"
	"fmt"
	"strings"
)

type coord struct{ x, y int }

func (c *coord) toString() string {
	return fmt.Sprintf("%d;%d", c.x, c.y)
}

func (c *coord) applyDirection(direction string) error {
	switch direction {
	case "^":
		c.y++
	case "v":
		c.y--
	case ">":
		c.x++
	case "<":
		c.x--
	default:
		return errors.New("nope")
	}
	return nil
}

func runStep(visited map[string]int, current *coord, direction string) {
	if current.applyDirection(direction) == nil {
		visited[current.toString()]++
	}
}

func solo(directions []string) int {
	var visited = make(map[string]int)
	var current = coord{}
	visited[current.toString()]++

	for _, direction := range directions {
		runStep(visited, &current, direction)
	}

	return len(visited)
}

func duo(directions []string) int {
	visited := make(map[string]int)
	santa := coord{}
	roboSanta := coord{}

	for index, direction := range directions {
		if index%2 == 0 {
			runStep(visited, &santa, direction)
		} else {
			runStep(visited, &roboSanta, direction)
		}
	}

	return len(visited)
}

func Day3(input string) (int, int) {
	directions := strings.Split(input, "")

	return solo(directions), duo(directions)
}
