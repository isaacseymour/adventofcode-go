package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
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

type Box struct {
	length int
	width  int
	height int
}

func (box *Box) sideAreas() (int, int, int) {
	return box.length * box.width, box.length * box.height, box.width * box.height
}

func (box *Box) smallestSide() int {
	// Box side lengths are sorted, so length < width < height, hence
	return box.length * box.width
}

func (box *Box) surfaceArea() int {
	x, y, z := box.sideAreas()
	return 2 * (x + y + z)
}

func (box *Box) ribbonLength() int {
	return (2 * (box.length + box.width)) + (box.length * box.width * box.height)
}

func parseBox(input string) (Box, error) {
	parts := strings.Split(input, "x")
	if len(parts) != 3 {
		return Box{}, errors.New(fmt.Sprintf("Invalid number of parts in %s: %d", input, len(parts)))
	}

	var lwh = make([]int, 3, 3)
	var err error
	for i, dimen := range parts {
		lwh[i], err = strconv.Atoi(dimen)
		if err != nil {
			return Box{}, errors.New(fmt.Sprintf("Error parsing dimension %s: %s", dimen, err))
		}
	}

	sort.Ints(lwh)
	return Box{lwh[0], lwh[1], lwh[2]}, nil
}

func Day2(input string) (int, int) {
	var totalArea int
	var totalRibbon int

	for _, line := range strings.Split(input, "\n") {
		box, err := parseBox(line)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error parsing %s: %s", line, err))
		}

		totalArea += box.surfaceArea() + box.smallestSide()
		totalRibbon += box.ribbonLength()
	}

	return totalArea, totalRibbon
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	strinput := string(input)

	day, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Day", day)

	var result string
	switch day {
	case 1:
		floor, firstBasement := Day1(strinput)
		result = fmt.Sprintf("End floor: %d; first enters basement at %d", floor, firstBasement)
	case 2:
		area, ribbon := Day2(strinput)
		result = fmt.Sprintf("Area: %d; ribbon: %d", area, ribbon)
	}

	fmt.Println(result)
}
