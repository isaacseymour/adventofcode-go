package adventofcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Instruction int

const (
	On Instruction = iota
	Off
	Toggle
)

type command struct {
	instruction Instruction
	startX      int
	startY      int
	endX        int
	endY        int
}

func (c *command) applyTo(grid map[int]map[int]bool) {
	for x := c.startX; x <= c.endX; x++ {
		for y := c.startY; y <= c.endY; y++ {
			if grid[x] == nil {
				grid[x] = make(map[int]bool)
			}

			switch c.instruction {
			case On:
				grid[x][y] = true
			case Off:
				grid[x][y] = false
			case Toggle:
				grid[x][y] = !grid[x][y]
			}
		}
	}
}

func (c *command) applyTo2(grid map[int]map[int]int) {
	for x := c.startX; x <= c.endX; x++ {
		for y := c.startY; y <= c.endY; y++ {
			if grid[x] == nil {
				grid[x] = make(map[int]int)
			}

			switch c.instruction {
			case On:
				grid[x][y]++
			case Off:
				if grid[x][y] > 0 {
					grid[x][y]--
				}
			case Toggle:
				grid[x][y] += 2
			}
		}
	}
}

func parseCoords(input string) (int, int, error) {
	coords := strings.Split(input, ",")

	if len(coords) != 2 {
		return -1, -1, errors.New(fmt.Sprintf("Invalid coordinates %s", input))
	}

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return -1, -1, err
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return -1, -1, err
	}

	return x, y, nil
}

func parseCommand(input string) (command, error) {
	tokens := strings.Split(input, " ")
	c := command{}

	if tokens[0] == "toggle" {
		tokens = tokens[1:]
		c.instruction = Toggle
	} else if tokens[0] == "turn" && tokens[1] == "on" {
		tokens = tokens[2:]
		c.instruction = On
	} else if tokens[0] == "turn" && tokens[1] == "off" {
		tokens = tokens[2:]
		c.instruction = Off
	} else {
		return c, errors.New(fmt.Sprintf("Unable to parse %s", input))
	}

	startX, startY, err := parseCoords(tokens[0])
	if err != nil {
		return c, err
	}
	c.startX = startX
	c.startY = startY

	tokens = tokens[1:]

	if tokens[0] != "through" {
		return c, errors.New("missing through!")
	}

	tokens = tokens[1:]

	endX, endY, err := parseCoords(tokens[0])

	if err != nil {
		return c, err
	}
	c.endX = endX
	c.endY = endY

	if len(tokens) != 1 {
		return c, errors.New(fmt.Sprintf("unexpected tokens at end: %s", strings.Join(tokens[1:], " ")))
	}

	return c, nil
}

func Day6(input string) (int, int) {
	grid1 := make(map[int]map[int]bool)
	grid2 := make(map[int]map[int]int)

	for _, line := range strings.Split(input, "\n") {
		c, err := parseCommand(line)
		if err == nil {
			c.applyTo(grid1)
			c.applyTo2(grid2)
		}
	}

	count := 0
	for _, ys := range grid1 {
		for _, status := range ys {
			if status {
				count++
			}
		}
	}

	total := 0
	for _, ys := range grid2 {
		for _, brightness := range ys {
			total += brightness
		}
	}

	return count, total
}
