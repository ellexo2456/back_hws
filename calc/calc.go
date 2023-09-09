package calc

import (
	"fmt"
	"strings"
)

func getDeepestNearestCols(expr string) string {
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
			return getDeepestNearestCols(expr[startIndex+1 : endIndex+1])
		}
	}

	return expr
}

func Calc(expression string) float64 {

	//var result float64
	currentExpr := expression
	for currentExpr != "" {
		if strings.Contains(currentExpr, "(") {
			currentExpr = getDeepestNearestCols(currentExpr)
		}
		fmt.Println(currentExpr)
		return 4
		//result += calculate(line)
		//input = input.replaceInColsWithResult()

	}

	return 5
}
