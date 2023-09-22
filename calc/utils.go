package calc

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type utilsMathComponents struct {
	memorizedNumber   string
	memorizedOperator string
	curGetNumber      string
	curGetOperator    string
	curNumber         string
	nextOperator      string
}

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

func numbersUpdate(components utilsMathComponents) (utilsMathComponents, error) {
	var err error
	if components.curNumber, err = makeOperation(components.curNumber, components.curGetNumber, components.curGetOperator); err != nil {
		return utilsMathComponents{}, err
	}

	if components.memorizedNumber != "" {
		if components.curNumber, err = makeOperation(components.memorizedNumber, components.curNumber, components.memorizedOperator); err != nil {
			return utilsMathComponents{}, err
		}
		components.memorizedNumber = ""
	}

	return components, nil
}

func operatorSwitch(components utilsMathComponents) (utilsMathComponents, error) {
	var err error

	switch components.nextOperator {
	case "+", "-":
		if components, err = numbersUpdate(components); err != nil {
			return utilsMathComponents{}, err
		}

	case "*", "/":
		if components.memorizedNumber != "" {
			if components.curNumber, err = makeOperation(components.curNumber, components.curGetNumber, components.curGetOperator); err != nil {
				return utilsMathComponents{}, err
			}

		} else {
			components.memorizedNumber = components.curNumber
			components.memorizedOperator = components.curGetOperator
			components.curNumber = components.curGetNumber
		}
	}

	return components, nil
}

func calculateLoop(expr string, components utilsMathComponents, operatorIndex int) (string, error) {
	var err error

	for expr != "" {
		components.curGetNumber, operatorIndex = getNumber(expr)
		if operatorIndex == -1 {
			return "", errors.New("error: invalid string")
		}

		if operatorIndex == len(expr) {
			if components, err = numbersUpdate(components); err != nil {
				return "", err
			}
			return components.curNumber, nil
		}

		components.nextOperator = string(expr[operatorIndex])
		expr = expr[operatorIndex+1:]

		if components, err = operatorSwitch(components); err != nil {
			return "", err
		}

		components.curGetOperator = components.nextOperator
	}

	return "", errors.New("error: incorrect expression")
}

func calculate(expr string) (string, error) {
	if len(expr) == 0 || len(expr) == 1 {
		return expr, nil
	}

	var components utilsMathComponents
	var err error
	var operatorIndex int
	if expr[0] == '-' {
		expr, components.curNumber, components.curGetOperator, operatorIndex, err = getFirstNumber(expr[1:])
		components.curNumber = "-" + components.curNumber
	} else {
		expr, components.curNumber, components.curGetOperator, operatorIndex, err = getFirstNumber(expr)
	}

	if err != nil {
		return "", errors.New("error: invalid string")
	}
	if expr == "" {
		return components.curNumber, nil
	}

	return calculateLoop(expr, components, operatorIndex)
}

func replaceInParenthesisWithResult(line string, inParenthesis string, result string) string {
	inParenthesis = "(" + inParenthesis + ")"
	return strings.Replace(line, inParenthesis, result, -1)
}
