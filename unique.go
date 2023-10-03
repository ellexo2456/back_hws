package main

import (
	"back_hws/unique"
	"fmt"
)

func main() {
	options := optionsInit()

	input, err := readInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if input, err = unique.Unique(input, options); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if err = writeOutput(input); err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
