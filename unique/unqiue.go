package unique

func Unique(lines []string) []string {
	var unique_lines []string
	unique_lines = append(unique_lines, lines[0])
	for i, line := range lines[1:] {
		if line != lines[i] {
			unique_lines = append(unique_lines, line)
		}
	}
	return unique_lines
}
