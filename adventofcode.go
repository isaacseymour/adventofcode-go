package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Challenge1(input string) int {
	var result int

	for _, char := range strings.Split(input, "") {
		if char == "(" {
			result++
		} else if char == ")" {
			result--
		} else {
			fmt.Print(fmt.Sprintf("ignoring invalid character %s", char))
		}
	}

	return result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print("Challenge 1: ", Challenge1(string(input)))
}
