package calc

import (
	"errors"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0, errors.New("error: empty input")
	}
	expression = strings.Join(strings.Fields(expression), "")

	var curResult string
	currentExpr := expression
	var err error
	for strings.Contains(expression, "(") {
		if currentExpr, err = getDeepestNearestParenthesis(expression); err != nil {
			return 0, err
		}

		if curResult, err = calculate(currentExpr); err != nil {
			return 0, err
		}

		expression = replaceInParenthesisWithResult(expression, currentExpr, curResult)
	}

	if curResult, err = calculate(expression); err != nil {
		return 0, err
	}

	var result float64
	if result, err = strconv.ParseFloat(curResult, 64); err != nil {
		return 0, err
	}
	return result, nil
}
