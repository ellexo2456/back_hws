package unique

import "strconv"

type Options struct {
	C *bool
	D *bool
	U *bool
}

func Unique(lines []string, options Options) []string {
	uniqueLines := []string{lines[0]}
	count := 1

	for i, line := range lines[1:] {

		if line != lines[i] {
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
