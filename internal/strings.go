package internal

func BreakLongLine(line string, maxLineLength int) []string {
	var lines []string

	for len(line) > maxLineLength {
		lines = append(lines, line[:maxLineLength])
		line = line[maxLineLength:]
	}

	lines = append(lines, line)

	return lines
}
