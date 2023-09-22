package main

import (
	"back_hws/unique"
	"log"
)

func main() {
	options := optionsInit()

	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	if input, err = unique.Unique(input, options); err != nil {
		log.Fatal(err)
	}

	if err = writeOutput(input); err != nil {
		log.Fatal(err)
	}
}
