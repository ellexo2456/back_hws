package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tp_go_course_second_sem_2023/unique"
)

func main() {
	//inpt := "I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik."

	var input_file *os.File
	var output_file *os.File
	var err error
	params := os.Args[1:]
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
