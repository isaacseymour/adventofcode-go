package adventofcode

import (
	"strconv"
	"strings"
)

func Day8(input string) (int, int) {
	totalChars, totalUnquotedChars, totalQuotedChars := 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		totalChars += len(line)
		unquoted, _ := strconv.Unquote(line)
		totalUnquotedChars += len(unquoted)
		totalQuotedChars += len(strconv.Quote(line))
	}

	return totalChars - totalUnquotedChars, totalQuotedChars - totalChars
}
