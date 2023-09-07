package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"tp_go_course_second_sem_2023/unique"
)

type Options struct {
	cFlag *bool
}

func main() {
	options := Options{
		cFlag: flag.Bool("c", false, "help message for flag n"),
	}

	flag.Parse()
	if *options.cFlag {
		fmt.Println("egeoom")
	}

	var input_file *os.File
	var output_file *os.File
	var err error
	params := flag.Args()
	switch len(params) {
	case 1:
		input_file, err = os.Open(params[0])
		defer input_file.Close()
	case 2:
		input_file, err = os.Open(params[0])
		output_file, err = os.Create(params[1])
		defer input_file.Close()
		defer output_file.Close()

	}
	if err != nil {
		log.Fatal(err)
	}

	var scanner *bufio.Scanner
	if input_file != nil {
		scanner = bufio.NewScanner(input_file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	input = unique.Unique(input)
	var writer *bufio.Writer
	if output_file != nil {
		writer = bufio.NewWriter(output_file)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	for _, i := range input {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
