package main

import (
	"back_hws/calc"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string
	input = scanner.Text()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	var result float64
	if result, err = calc.Calc(input); err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(result)
}
