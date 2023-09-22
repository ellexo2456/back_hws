package main

import (
	"back_hws/unique"
	"flag"
	"log"
)

func main() {
	options := optionsInit()
	flag.Parse()

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
