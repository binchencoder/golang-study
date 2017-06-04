package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	wordfreq := make(map[string]int)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("The input was: %s\n", input)
	}

	s := strings.Split(input, " ")
	for _, word := range s {
		fmt.Println(word)

		wordfreq[word]++
	}

	// input := bufio.NewScanner(os.Stdin)
	// input.Split(bufio.ScanWords)
	// for input.Scan() {
	// 	line := input.Text()

	// 	wordfreq[line]++
	// 	fmt.Printf("Wordfreq map: %v+\n", wordfreq)
	// }

	for str, count := range wordfreq {
		fmt.Printf("String of %s frequency %d\r\n", str, count)

		// if count, ok := wordfreq[line]; ok {
		// 	wordfreq[line]++
		// 	fmt.Printf("String of %s frequency %d\n", line, count)
		// }
	}
}
