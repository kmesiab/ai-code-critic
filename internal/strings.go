package internal

import "strings"

func BreakLongLine(line string, maxLineLength int) []string {
	var lines []string

	for len(line) > maxLineLength {
		lines = append(lines, line[:maxLineLength])
		line = line[maxLineLength:]
	}

	lines = append(lines, line)

	return lines
}

func ShortenLongLines(input string) string {
	var output string

	for _, line := range strings.Split(input, "\n") {
		output += strings.Join(BreakLongLine(line, 80), "\n")
		output += "\n"
	}

	return output
}
