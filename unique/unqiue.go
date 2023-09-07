package unique

import "strconv"

type Options struct {
	C *bool
}

func Unique(lines []string, options Options) []string {
	uniqueLines := []string{lines[0]}
	count := 1

	for i, line := range lines[1:] {

		if line != lines[i] {
			if *options.C {
				uniqueLines[len(uniqueLines)-1] = strconv.Itoa(count) + " " + uniqueLines[len(uniqueLines)-1]
				count = 1
			}

			uniqueLines = append(uniqueLines, line)
			continue
		}

		count++
	}
	if *options.C {
		uniqueLines[len(uniqueLines)-1] = strconv.Itoa(count) + " " + uniqueLines[len(uniqueLines)-1]
	}

	return uniqueLines
}
