package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			log.Print(err)
			continue
		}
		wordfreq(f, counts)
		f.Close()
	}
	for word, n := range counts {
		fmt.Printf("%s\t%d\n", word, n)
	}
}

func wordfreq(f io.Reader, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
