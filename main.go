package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	//calling the count func to count the number of words (or lines)
	// received from stdin and printing it out
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Define a counter
	wc := 0

	// scan the input and increment the counter
	for scanner.Scan() {
		wc++
	}

	//return the total count
	return wc
}
