package main

import (
	"bufio"
	"go-tasks/ciphers"
	"log"
	"os"
	"strings"
)

func ReadInput(file string) string {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Could not open an input file")
	}

	scanner := bufio.NewScanner(f)

	input := ""

	for scanner.Scan() {
		input += scanner.Text()
	}

	return input
}

func WriteInput(input *string) {
	output, err := os.Create("output.txt")

	defer output.Close()

	if err != nil {
		panic("Could not create an output file")
	}

	_, err2 := output.WriteString(*input)
	if err2 != nil {
		panic("Error writing data")
	}
}

func main() {
	txt := ReadInput("input.txt")
	args := strings.Split(txt, ",")

	cipher := ciphers.Cipher{}
	res := cipher.Handle(&args[0], args[1])

	WriteInput(res)
}
