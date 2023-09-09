package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"tp_go_course_second_sem_2023/unique"
)

func main() {
	options := OptionsInit()
	flag.Parse()

	inputFile, inputError, outputFile, outputError := FileInit()

	if inputError != nil {
		log.Fatal(inputError)
	}
	if outputError != nil {
		log.Fatal(outputError)
	}

	var scanner *bufio.Scanner
	if inputFile != nil {
		scanner = bufio.NewScanner(inputFile)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
		defer inputFile.Close()
	}

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	input, err = unique.Unique(input, options)
	if err != nil {
		log.Fatal(err)
	}

	var writer *bufio.Writer
	if outputFile != nil {
		writer = bufio.NewWriter(outputFile)
	} else {
		writer = bufio.NewWriter(os.Stdout)
		defer outputFile.Close()
	}

	for _, i := range input {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
