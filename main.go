package main

import (
	aoc "./adventofcode"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	strinput := string(input)

	day, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Day", day)

	var result string
	switch day {
	case 1:
		floor, firstBasement := aoc.Day1(strinput)
		result = fmt.Sprintf("End floor: %d; first enters basement at %d", floor, firstBasement)
	case 2:
		area, ribbon := aoc.Day2(strinput)
		result = fmt.Sprintf("Area: %d; ribbon: %d", area, ribbon)
	case 3:
		solo, duo := aoc.Day3(strinput)
		result = fmt.Sprintf("First year: %d; second year: %d", solo, duo)
	case 4:
		five, six := aoc.Day4(strinput)
		result = fmt.Sprintf("Five: %d; six: %d", five, six)
	case 5:
		nice1, nice2 := aoc.Day5(strinput)
		result = fmt.Sprintf("Nice1: %d; nice2: %d", nice1, nice2)
	case 6:
		count, brightness := aoc.Day6(strinput)
		result = fmt.Sprintf("On: %d, brightness: %d", count, brightness)
	case 7:
		output := aoc.Day7(strinput)
		result = fmt.Sprintf("Result! %d", output)
	}

	fmt.Println(result)
}
