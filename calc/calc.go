package calc

import (
	"errors"
	"fmt"
	"strings"
)

func getDeepestNearestParenthesis(expr string) (string, error) {
	var startIndex int
	if startIndex = strings.Index(expr, "("); startIndex == -1 {
		return expr, nil
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
			return getDeepestNearestParenthesis(expr[startIndex+1 : endIndex+1])
		}
	}

	return expr, errors.New("incorrect count of parenthesis")
}

func Calc(expression string) (float64, error) {

	//var result float64
	currentExpr := expression
	var err error
	for currentExpr != "" {
		if strings.Contains(currentExpr, "(") {
			if currentExpr, err = getDeepestNearestParenthesis(currentExpr); err != nil {
				return 0, err
			}
		}
		fmt.Println(currentExpr)
		return 4, nil
		//result += calculate(line)
		//input = input.replaceInColsWithResult()

	}

	return 5, nil
}
