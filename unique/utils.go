package unique

import (
	"errors"
	"strconv"
	"strings"
)

func argumentsCheck(lines []string, options Options) ([]string, error) {
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

func prepareToCompare(curLine string, prevLine string, options Options) (string, string) {

	if *options.I {
		curLine, prevLine = strings.ToLower(curLine), strings.ToLower(prevLine)
	}
	if *options.F != 0 {
		curLine, prevLine = fieldCutter(strings.Fields(curLine), *options.F), fieldCutter(strings.Fields(prevLine), *options.F)
	}
	if *options.S != 0 {
		curLine, prevLine = runeCutter([]rune(curLine), *options.S), runeCutter([]rune(prevLine), *options.S)
	}

	return curLine, prevLine
}

func formatLinesSlice(options Options, slice []string, count int) []string {
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
