package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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

func calculate(expr string) (float64, error) {
	if len(expr) == 0 || len(expr) == 0 {
		return strconv.ParseFloat(expr, 64)
	}

	var result float64
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
			if result, err = strconv.ParseFloat(curNumber, 64); err != nil {
				return 0, nil
			}
			return result, nil
		}

		curGetOperator = string(expr[operatorIndex+1])
		expr = expr[operatorIndex+2:]
	} else {
		curNumber, operatorIndex = getNumber(expr)

		if operatorIndex == len(expr) {
			if result, err = strconv.ParseFloat(curNumber, 64); err != nil {
				return 0, nil
			}
			return result, nil
		}

		curGetOperator = string(expr[operatorIndex])
		expr = expr[operatorIndex+1:]
	}

	for expr != "" {
		curGetNumber, operatorIndex = getNumber(expr)

		if operatorIndex == len(expr) {
			if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
				return 0, err
			}

			if memorizedNumber != "" {
				if curNumber, err = makeOperation(memorizedNumber, curNumber, memorizedOperator); err != nil {
					return 0, err
				}
			}

			if result, err = strconv.ParseFloat(curNumber, 64); err != nil {
				return 0, nil
			}
			return result, nil
		}

		nextOperator = string(expr[operatorIndex])
		expr = expr[operatorIndex+1:]

		switch nextOperator {
		case "+", "-":
			if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
				return 0, err
			}

			if memorizedNumber != "" {
				if curNumber, err = makeOperation(memorizedNumber, curNumber, memorizedOperator); err != nil {
					return 0, err
				}
				memorizedNumber = ""
			}

		case "*", "/":
			if memorizedNumber != "" {
				if curNumber, err = makeOperation(curNumber, curGetNumber, curGetOperator); err != nil {
					return 0, err
				}

			} else {
				memorizedNumber = curNumber
				memorizedOperator = curGetOperator
				curNumber = curGetNumber
			}
		}

		curGetOperator = nextOperator
	}

	return 0, errors.New("error: incorrect expression")
}

func Calc(expression string) (float64, error) {

	var result float64
	var curResult float64
	currentExpr := expression
	var err error
	for currentExpr != "" {
		if strings.Contains(currentExpr, "(") {
			if currentExpr, err = getDeepestNearestParenthesis(currentExpr); err != nil {
				return 0, err
			}
		}

		if curResult, err = calculate(currentExpr); err != nil {
			return 0, err
		}
		result += curResult
		fmt.Println(result)
		break
		//input = input.replaceInColsWithResult()

	}

	return 5, nil
}
