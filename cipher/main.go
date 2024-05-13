package main

import (
	"bufio"
	"fmt"
	"go-tasks/ciphers"
	"log"
	"os"
	"strings"
)

func readInput(file string) (string, error) {
	f, err := os.Open(file)

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(f)

	input := ""

	for scanner.Scan() {
		input += scanner.Text()
	}

	return input, nil
}

func writeInput(input string) (int, error) {
	output, err := os.Create("output.txt")

	defer output.Close()

	res, err := output.WriteString(input)
	return res, err
}

func main() {
	text, err := readInput("input.txt")

	if err != nil {
		log.Fatal("Could not read the file")
	}

	args := strings.Split(text, ",")

	cipher := ciphers.New(args[0], args[1])

	res, err := cipher.Handle()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	_, err = writeInput(res)

	if err != nil {
		log.Fatal("Could not write input")
	}
}
