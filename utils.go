package main

import (
	"back_hws/unique"
	"flag"
	"os"
)

func OptionsInit() unique.Options {
	return unique.Options{
		C: flag.Bool("c", false, "count the number of line appearances"),
		D: flag.Bool("d", false, "only duplicated"),
		U: flag.Bool("u", false, "only unique"),
		I: flag.Bool("i", false, "ignore case"),
		F: flag.Int("f", 0, "ignore first num fields"),
		S: flag.Int("s", 0, "ignore first num chars"),
	}
}

func FileInit() (*os.File, error, *os.File, error) {
	params := flag.Args()

	switch len(flag.Args()) {
	case 1:
		inputFile, err := os.Open(params[0])
		return inputFile, err, nil, nil
	case 2:
		inputFile, inputErr := os.Open(params[0])
		outputFile, outputErr := os.Create(params[1])
		return inputFile, inputErr, outputFile, outputErr
	default:
		return nil, nil, nil, nil
	}
}
