package unique

import (
	"errors"
	"strconv"
	"strings"
)

type pair struct {
	values, offset interface{}
}

type a struct {
	values []string
	offset int
}

func ArgumentsCheck(lines []string, options Options) ([]string, error) {
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

	return lines, nil
}

func fieldCutter(fields []string, count int) string {
	if count > len(fields) {
		return ""
	} else {
		return strings.Join(fields[count:], " ")
	}
}

func runeCutter(fields []rune, count int) string {
	if count > len(fields) {
		return ""
	} else {
		return string(fields[count:])
	}
}

func PrepareToCompare(curLine string, prevLine string, options Options) (string, string) {
	if *options.I {
		return strings.ToLower(curLine), strings.ToLower(prevLine)
	}
	if *options.F != 0 {
		//lineToCompareFields := strings.Fields(curLine)
		//linesIFields := strings.Fields(prevLine)
		//
		//if *options.F > len(lineToCompareFields) {
		//	curLine = ""
		//} else {
		//	curLine = strings.Join(lineToCompareFields[*options.F:], " ")
		//}
		//
		//if *options.F > len(linesIFields) {
		//	prevLine = ""
		//} else {
		//	prevLine = strings.Join(linesIFields[*options.F:], " ")
		//}
		return fieldCutter(strings.Fields(curLine), *options.F), fieldCutter(strings.Fields(prevLine), *options.F)
	}
	if *options.S != 0 {
		//lineToCompareRunes := []rune(curLine)
		//linesIRunes := []rune(prevLine)
		//
		//if *options.S > len(lineToCompareRunes) {
		//	curLine = ""
		//} else {
		//	curLine = string(lineToCompareRunes[*options.S:])
		//}
		//
		//if *options.S > len(linesIRunes) {
		//	prevLine = ""
		//} else {
		//	prevLine = string(linesIRunes[*options.S:])
		//}

		return runeCutter([]rune(curLine), *options.F), runeCutter([]rune(curLine), *options.F)

	}

	return curLine, prevLine
}

func FormatLinesSlice(options Options, slice []string, count int) []string {
	if *options.C {
		slice[len(slice)-1] = strconv.Itoa(count) + " " + slice[len(slice)-1]
	} else if *options.D {
		if count == 1 {
			slice = slice[:len(slice)-1]
		}
	} else if *options.U {
		if count > 1 {
			slice = slice[:len(slice)-1]
		}
	}

	return slice
}
