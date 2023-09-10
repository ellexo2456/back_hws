package calc

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func getDeepestNearestParenthesis(expr string) (string, error) {
	if strings.Count(expr, "(") != strings.Count(expr, ")") {
		return "", errors.New("incorrect count of parenthesis")
	}

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

func innerGetNumber(expr string) (string, int) {
	var number string

	for index, symbolCode := range expr {
		if unicode.IsNumber(symbolCode) || symbolCode == '.' || symbolCode == ',' {
			number += string(symbolCode)
		} else {
			if utf8.RuneCountInString(number) == 0 {
				return "", -1
			}

			return number, index
		}
	}

	return number, len(expr)
}

func getNumber(expr string) (string, int) {
	if utf8.RuneCountInString(expr) == 0 {
		return "", -1
	}

	var number string
	var index int
	if expr[0] == '-' {
		if number, index = innerGetNumber(expr[1:]); index == -1 {
			return number, index
		} else {
			return "-" + number, index + 1
		}
	}
	return innerGetNumber(expr)
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

func getFirstNumber(expr string) (string, string, string, int, error) {
	curNumber, operatorIndex := getNumber(expr)
	if operatorIndex == -1 {
		return "", "", "", 0, errors.New("error: invalid string")
	}

	if operatorIndex == len(expr) {
		return "", curNumber, "", 0, nil
	}

	return expr[operatorIndex+1:], curNumber, string(expr[operatorIndex]), operatorIndex, nil
}

func NumbersUpdate(curNumber string, memorizedNumber string, curGetNumber string, curGetOperator string, memorizedOperator string) (string, string, error) {
	var err error
	if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
		return "", "", err
	}

	if memorizedNumber != "" {
		if curNumber, err = makeOperation(memorizedNumber, curNumber, memorizedOperator); err != nil {
			return "", "", err
		}
		memorizedNumber = ""
	}

	return curNumber, memorizedNumber, nil
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
		expr, curNumber, curGetOperator, operatorIndex, err = getFirstNumber(expr[1:])
		curNumber = "-" + curNumber
	} else {
		expr, curNumber, curGetOperator, operatorIndex, err = getFirstNumber(expr)
	}

	if err != nil {
		return "", errors.New("error: invalid string")
	}
	if expr == "" {
		return curNumber, nil
	}

	for expr != "" {
		curGetNumber, operatorIndex = getNumber(expr)
		if operatorIndex == -1 {
			return "", errors.New("error: invalid string")
		}

		if operatorIndex == len(expr) {
			if curNumber, memorizedNumber, err = NumbersUpdate(curNumber, memorizedNumber, curGetNumber, curGetOperator, memorizedOperator); err != nil {
				return "", err
			}
			return curNumber, nil
		}

		nextOperator = string(expr[operatorIndex])
		expr = expr[operatorIndex+1:]

		switch nextOperator {
		case "+", "-":
			if curNumber, memorizedNumber, err = NumbersUpdate(curNumber, memorizedNumber, curGetNumber, curGetOperator, memorizedOperator); err != nil {
				return "", err
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

func replaceInParenthesisWithResult(line string, inParenthesis string, result string) string {
	inParenthesis = "(" + inParenthesis + ")"
	return strings.Replace(line, inParenthesis, result, -1)
}
