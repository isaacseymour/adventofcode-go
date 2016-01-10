package adventofcode

import (
	"strings"
)

func hasThreeVowels(x string) bool {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0

	for _, char := range x {
		if vowels[char] {
			count++
		}

		if count >= 3 {
			return true
		}
	}

	return false
}

func containsDouble(x string) bool {
	last := x[0]
	for i := 1; i < len(x); i++ {
		if last == x[i] {
			return true
		}
		last = x[i]
	}

	return false
}

func containsBanned(x string) bool {
	banned := []string{"ab", "cd", "pq", "xy"}
	for _, b := range banned {
		if strings.Contains(x, b) {
			return true
		}
	}
	return false
}

func nice(x string) bool {
	return hasThreeVowels(x) && containsDouble(x) && !containsBanned(x)
}

func twoPairs(x string) bool {
	for i := 1; i < len(x); i++ {
		pair := x[i-1 : i+1]
		if strings.Contains(x[i+1:], pair) {
			return true
		}
	}

	return false
}

func surroundedPair(x string) bool {
	if len(x) < 3 {
		return false
	}

	for i := 2; i < len(x); i++ {
		if x[i-2] == x[i] {
			return true
		}
	}

	return false
}

func nice2(x string) bool {
	return twoPairs(x) && surroundedPair(x)
}

func Day5(input string) (int, int) {
	count1 := 0
	count2 := 0
	for _, x := range strings.Split(input, "\n") {
		if nice(x) {
			count1++
		}

		if nice2(x) {
			count2++
		}
	}
	return count1, count2
}
