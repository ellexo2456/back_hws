package calc

import "strings"

func getDeepestCols(expr string) string {
	var startIndex int
	if startIndex = strings.Index(expr, "("); startIndex == -1 {
		return expr
	}

	count := 1
	for endIndex, symbolCode := range expr[startIndex+1:] {
		switch string(symbolCode) {
		case "(":
			count++
		case ")":
			count--
		}

		if count == 0 {
			return getDeepestCols(expr[startIndex:endIndex])
		}
	}

	return expr
}

func Calc(expression string) float64 {

	var result float64
	currentExpr := expression
	for currentExpr == "" {
		if strings.Contains(currentExpr, "(") {
			currentExpr = getDeepestCols()
		}
		result += calculate(line)
		input = input.replaceInColsWithResult()

	}

}
