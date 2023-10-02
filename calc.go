package main

import (
	"back_hws/calc"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string
	input = scanner.Text()

	err := scanner.Err()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var result float64
	if result, err = calc.Calc(input); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(result)
}
