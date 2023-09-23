package main

import (
	"back_hws/unique"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func optionsInit() unique.Options {
	c := flag.Bool("c", false, "count the number of line appearances")
	d := flag.Bool("d", false, "only duplicated")
	u := flag.Bool("u", false, "only unique")
	i := flag.Bool("i", false, "ignore case")
	f := flag.Int("f", 0, "ignore first num fields")
	s := flag.Int("s", 0, "ignore first num chars")

	flag.Parse()

	return unique.Options{
		C: *c,
		D: *d,
		U: *u,
		I: *i,
		F: *f,
		S: *s,
	}
}

func closeAndReport(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error; ", err)
	}
}

func readInput() ([]string, error) {
	params := flag.Args()
	var inputFile *os.File
	var err error
	var scanner *bufio.Scanner

	if len(params) >= 1 {
		if inputFile, err = os.Open(params[0]); err != nil {
			return nil, err
		}
		defer closeAndReport(inputFile)
		scanner = bufio.NewScanner(inputFile)

	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func writeOutput(input []string) error {
	params := flag.Args()
	var outputFile *os.File
	var err error
	var writer *bufio.Writer

	if len(params) == 2 {
		if outputFile, err = os.Create(params[1]); err != nil {
			return err
		}
		closeAndReport(outputFile)
		writer = bufio.NewWriter(outputFile)

	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	for _, i := range input {
		if _, err := fmt.Fprintln(writer, i); err != nil {
			return err
		}
	}
	if err = writer.Flush(); err != nil {
		return err
	}

	return nil
}
