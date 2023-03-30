package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	words := flag.Bool("w", false, "count words")
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *words, *lines, *bytes))
}

func count(r io.Reader, words bool, lines bool, bytes bool) int {
	scanner := bufio.NewScanner(r)
	wc := 0

	if words {
		scanner.Split(bufio.ScanWords)
	} else if lines {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanBytes)
	}

	for scanner.Scan() {
		wc += 1
	}
	return wc
}
