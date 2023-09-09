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
		//lineToCompare := line
		//
		//if *options.I {
		//	lineToCompare = strings.ToLower(lineToCompare)
		//	lines[i] = strings.ToLower(lines[i])
		//}
		//if *options.F != 0 {
		//	lineToCompareFields := strings.Fields(lineToCompare)
		//	linesIFields := strings.Fields(lines[i])
		//
		//	if *options.F > len(lineToCompareFields) {
		//		lineToCompare = ""
		//	} else {
		//		lineToCompare = strings.Join(lineToCompareFields[*options.F:], " ")
		//	}
		//
		//	if *options.F > len(linesIFields) {
		//		lines[i] = ""
		//	} else {
		//		lines[i] = strings.Join(linesIFields[*options.F:], " ")
		//	}
		//}
		//if *options.S != 0 {
		//	lineToCompareRunes := []rune(lineToCompare)
		//	linesIRunes := []rune(lines[i])
		//
		//	if *options.S > len(lineToCompareRunes) {
		//		lineToCompare = ""
		//	} else {
		//		lineToCompare = string(lineToCompareRunes[*options.S:])
		//	}
		//
		//	if *options.S > len(linesIRunes) {
		//		lines[i] = ""
		//	} else {
		//		lines[i] = string(linesIRunes[*options.S:])
		//	}
		//}
		curLine, prevLine := PrepareToCompare(line, lines[i], options)
		if curLine != prevLine {
			//if *options.C {
			//	uniqueLines[len(uniqueLines)-1] = strconv.Itoa(count) + " " + uniqueLines[len(uniqueLines)-1]
			//} else if *options.D {
			//	if count == 1 {
			//		uniqueLines = uniqueLines[:len(uniqueLines)-1]
			//	}
			//} else if *options.U {
			//	if count > 1 {
			//		uniqueLines = uniqueLines[:len(uniqueLines)-1]
			//	}
			//}

			uniqueLines = FormatLinesSlice(options, uniqueLines, count)

			uniqueLines = append(uniqueLines, line)
			count = 1

			continue
		}

		count++
	}

	//if *options.C {
	//	uniqueLines[len(uniqueLines)-1] = strconv.Itoa(count) + " " + uniqueLines[len(uniqueLines)-1]
	//} else if *options.D {
	//	if count == 1 {
	//		uniqueLines = uniqueLines[:len(uniqueLines)-1]
	//	}
	//} else if *options.U {
	//	if count > 1 {
	//		uniqueLines = uniqueLines[:len(uniqueLines)-1]
	//	}
	//}

	return FormatLinesSlice(options, uniqueLines, count), nil
}
