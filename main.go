package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func writeHeader(file *os.File) {
	fmt.Fprintln(file, "#include<stdio.h>")
	fmt.Fprintln(file, "#include<stdlib.h>")
	fmt.Fprintln(file, "int main() {")
	fmt.Fprintln(file, "char tape[30000] = {0};")
	fmt.Fprintln(file, "char *p = tape;")
}

func writeFooter(file *os.File) {
	fmt.Fprintln(file, "return 0;")
	fmt.Fprintln(file, "}")
}

func writeCommands(input *os.File, output *os.File) {
	in, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	for _, char := range in {
		switch char {
		case '>':
			fmt.Fprintln(output, "++p;")
		case '<':
			fmt.Fprintln(output, "--p;")
		case '+':
			fmt.Fprintln(output, "++*p;")
		case '-':
			fmt.Fprintln(output, "--*p;")
		case '.':
			fmt.Fprintln(output, "putchar(*p);")
		case ',':
			fmt.Fprintln(output, "*p = getchar();")
		case '[':
			fmt.Fprintln(output, "while(*p){")
		case ']':
			fmt.Fprintln(output, "}")
		default:
			continue
		}
	}
}

func main() {
	var inputFileName string
	var outputFileName string
	flag.StringVar(&inputFileName, "i", "main.bf", "specify input file")
	flag.StringVar(&outputFileName, "o", inputFileName[:len(inputFileName)-3]+".c", "outputfile")
	flag.Parse()

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	outputFile, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	writeHeader(outputFile)
	writeCommands(inputFile, outputFile)
	writeFooter(outputFile)

	defer inputFile.Close()
	defer outputFile.Close()

}
