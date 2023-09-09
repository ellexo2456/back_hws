package unique

type Options struct {
	C *bool
	D *bool
	U *bool
	I *bool
	F *int
	S *int
}

func boolToInt(condition bool) int {
	if condition {
		return 1
	} else {
		return 0
	}
}

func Unique(lines []string, options Options) ([]string, error) {
	if _, err := ArgumentsCheck(lines, options); err != nil {
		return nil, err
	}

	if len(lines) <= 1 {
		return lines, nil
	}
	uniqueLines := []string{lines[0]}
	count := 1

	for i, line := range lines[1:] {
		curLine, prevLine := PrepareToCompare(line, lines[i], options)

		if curLine != prevLine {
			uniqueLines = FormatLinesSlice(options, uniqueLines, count)

			uniqueLines = append(uniqueLines, line)
			count = 1

			continue
		}

		count++
	}

	return FormatLinesSlice(options, uniqueLines, count), nil
}
