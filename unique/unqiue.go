package unique

import (
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

func Unique(lines []string, options Options) []string {
	if lines == nil {
		return nil
	}

	uniqueLines := []string{lines[0]}
	count := 1
	//strings.Join(strings.Fields(lines[0])[1:], " ")
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

	return uniqueLines
}
