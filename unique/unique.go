package unique

import (
	"errors"
	"strconv"
	"strings"
)

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
	if lines == nil {
		return nil, errors.New("Empty input")
	}
	if options == (Options{}) {
		return nil, errors.New("Empty options")
	}

	if *options.F < 0 {
		return nil, errors.New("Flag f must be non-negative")
	}
	if *options.S < 0 {
		return nil, errors.New("Flag s must be non-negative")
	}
	if (boolToInt(*options.C) + boolToInt(*options.D) + boolToInt(*options.U)) > 1 {
		return nil, errors.New("You`re can`t use flags c,d and u together")
	}

	uniqueLines := []string{lines[0]}
	count := 1

	for i, line := range lines[1:] {
		lineToCompare := line

		if *options.I {
			lineToCompare = strings.ToLower(lineToCompare)
			lines[i] = strings.ToLower(lines[i])
		}
		if *options.F != 0 {
			lineToCompareFields := strings.Fields(lineToCompare)
			linesIFields := strings.Fields(lines[i])

			if *options.F > len(lineToCompareFields) {
				lineToCompare = ""
			} else {
				lineToCompare = strings.Join(lineToCompareFields[*options.F:], " ")
			}

			if *options.F > len(linesIFields) {
				lines[i] = ""
			} else {
				lines[i] = strings.Join(linesIFields[*options.F:], " ")
			}
		}
		if *options.S != 0 {
			lineToCompareRunes := []rune(lineToCompare)
			linesIRunes := []rune(lines[i])

			if *options.S > len(lineToCompareRunes) {
				lineToCompare = ""
			} else {
				lineToCompare = string(lineToCompareRunes[*options.S:])
			}

			if *options.S > len(linesIRunes) {
				lines[i] = ""
			} else {
				lines[i] = string(linesIRunes[*options.S:])
			}
		}

		if lineToCompare != lines[i] {
			if *options.C {
				uniqueLines[len(uniqueLines)-1] = strconv.Itoa(count) + " " + uniqueLines[len(uniqueLines)-1]
			} else if *options.D {
				if count == 1 {
					uniqueLines = uniqueLines[:len(uniqueLines)-1]
				}
			} else if *options.U {
				if count > 1 {
					uniqueLines = uniqueLines[:len(uniqueLines)-1]
				}
			}

			uniqueLines = append(uniqueLines, line)
			count = 1

			continue
		}

		count++
	}

	if *options.C {
		uniqueLines[len(uniqueLines)-1] = strconv.Itoa(count) + " " + uniqueLines[len(uniqueLines)-1]
	} else if *options.D {
		if count == 1 {
			uniqueLines = uniqueLines[:len(uniqueLines)-1]
		}
	} else if *options.U {
		if count > 1 {
			uniqueLines = uniqueLines[:len(uniqueLines)-1]
		}
	}

	return uniqueLines, nil
}
