package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		counts := make(map[string]int)
		countLines(f, counts)
		f.Close()

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%#v\t%s\n", n, line, arg)
			}
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
