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
	b := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	//calling the count func to count the number of words (or lines)
	// received from stdin and printing it out
	fmt.Println(count(os.Stdin, *lines, *b))
	s := "word1 word2 word3 word4\n"
	fmt.Println(len(s))
}

func count(r io.Reader, countLines, countBytes bool) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
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
