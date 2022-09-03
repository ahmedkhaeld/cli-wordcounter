package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//calling the count func to count the number of words
	// received from stdin and printing it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {

	// A scanner is used to read text from a Reader (such as files)
	// Define the scanner split type to words (default is split by lines)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	// Define a counter
	wc := 0

	// scan the input and increment the counter
	for scanner.Scan() {
		wc++
	}

	return wc
}
