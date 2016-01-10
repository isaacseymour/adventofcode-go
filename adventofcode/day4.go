package adventofcode

import (
	"crypto/md5"
	"fmt"
)

func runHash(secret string, guess int) [md5.Size]byte {
	input := fmt.Sprintf("%s%d", secret, guess)
	return md5.Sum([]byte(input))
}

func fiveZeroes(hash [md5.Size]byte) bool {
	return hash[0] == 0 && hash[1] == 0 && hash[2] < 16
}

func sixZeroes(hash [md5.Size]byte) bool {
	return hash[0] == 0 && hash[1] == 0 && hash[2] == 0
}

func Day4(input string) (int, int) {
	var first int

	for i := 0; ; i++ {
		hash := runHash(input, i)

		if first == 0 && fiveZeroes(hash) {
			first = i
		}

		if sixZeroes(hash) {
			return first, i
		}
	}
}
