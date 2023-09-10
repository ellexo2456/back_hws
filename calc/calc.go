package calc

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
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
			return getDeepestNearestParenthesis(expr[startIndex+1 : startIndex+endIndex+1])
		}
	}

	return expr, errors.New("incorrect count of parenthesis")
}

func getNumber(expr string) (string, int) {
	var number string
	if expr[0] == '-' {
		number += "-"
		expr = expr[1:]
	}

	for index, symbolCode := range expr {
		if unicode.IsNumber(symbolCode) || symbolCode == '.' || symbolCode == ',' {
			number += string(symbolCode)
		} else {
			if number[0] == '-' {
				return number, index + 1
			}
			return number, index
		}
	}

	if number[0] == '-' {
		return number, len(expr) + 1
	}
	return number, len(expr)
}

func makeOperation(left string, right string, operator string) (string, error) {
	var leftNumber, rightNumber float64
	var err error
	if leftNumber, err = strconv.ParseFloat(left, 64); err != nil {
		return "", err
	}
	if rightNumber, err = strconv.ParseFloat(right, 64); err != nil {
		return "", err
	}

	var result float64
	switch operator {
	case "+":
		result = leftNumber + rightNumber
	case "-":
		result = leftNumber - rightNumber
	case "*":
		result = leftNumber * rightNumber
	case "/":
		result = leftNumber / rightNumber
	default:
		return "", errors.New("Error: incorrect operator")
	}

	return strconv.FormatFloat(result, 'f', -1, 64), nil
}

func calculate(expr string) (string, error) {
	if len(expr) == 0 || len(expr) == 1 {
		return expr, nil
	}

	var memorizedNumber, memorizedOperator string
	var curGetNumber, curGetOperator string
	var curNumber string
	var nextOperator string

	var err error
	var operatorIndex int

	if expr[0] == '-' {
		curNumber, operatorIndex = getNumber(expr[1:])
		curNumber = "-" + curNumber

		if operatorIndex == len(expr[1:]) {
			return curNumber, nil
		}

		curGetOperator = string(expr[operatorIndex+1])
		expr = expr[operatorIndex+2:]
	} else {
		curNumber, operatorIndex = getNumber(expr)

		if operatorIndex == len(expr) {
			return curNumber, nil
		}

		curGetOperator = string(expr[operatorIndex])
		expr = expr[operatorIndex+1:]
	}

	for expr != "" {
		curGetNumber, operatorIndex = getNumber(expr)

		if operatorIndex == len(expr) {
			if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
				return "", err
			}

			if memorizedNumber != "" {
				if curNumber, err = makeOperation(memorizedNumber, curNumber, memorizedOperator); err != nil {
					return "", err
				}
			}

			return curNumber, nil
		}

		nextOperator = string(expr[operatorIndex])
		expr = expr[operatorIndex+1:]

		switch nextOperator {
		case "+", "-":
			if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
				return "", err
			}

			if memorizedNumber != "" {
				if curNumber, err = makeOperation(memorizedNumber, curNumber, memorizedOperator); err != nil {
					return "", err
				}
				memorizedNumber = ""
			}

		case "*", "/":
			if memorizedNumber != "" {
				if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
					return "", err
				}

			} else {
				memorizedNumber = curNumber
				memorizedOperator = curGetOperator
				curNumber = curGetNumber
			}
		}

		curGetOperator = nextOperator
	}

	return "", errors.New("error: incorrect expression")
}

func isOneNumber(line string) bool {
	if utf8.RuneCountInString(line) == 1 && unicode.IsNumber(rune(line[0])) {
		return true
	}
	if line[0] == '-' {
		line = line[1:]
	}
	for _, symbolCode := range line {
		if !unicode.IsNumber(symbolCode) && symbolCode != '.' && symbolCode != ',' {
			return false
		}
	}

	return true
}

func replaceInParenthesisWithResult(line string, inParenthesis string, result string) string {
	inParenthesis = "(" + inParenthesis + ")"
	return strings.Replace(line, inParenthesis, result, -1)
}

func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0, errors.New("error: empty input")
	}
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
