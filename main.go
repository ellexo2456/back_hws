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
	options := unique.Options{
		C: flag.Bool("c", false, "count the number of line appearances"),
		D: flag.Bool("d", false, "only duplicated"),
		U: flag.Bool("u", false, "only unique"),
		I: flag.Bool("i", false, "ignore case"),
		F: flag.Int("f", 0, "ignore first num fields"),
		S: flag.Int("s", 0, "ignore first num chars"),
	}
	flag.Parse()

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
		log.Fatal(err)
	}

	input, err = unique.Unique(input, options)
	if err != nil {
		log.Fatal(err)
	}

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
