package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	for _, arg := range file {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
			continue
		}
		wordfreq(f, counts)
		f.Close()
	}
	for line, n := range counts {
		fmt.Printf("%s\t%d\n", line, n)
	}
}

func wordfreq(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
