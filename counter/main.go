package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	f, err := os.Open("/home/tatyana/gocode/src/tanya/hello/main.go")
	if err != nil {
		log.Fatal(err)
	}
	countLines(f, counts)
	f.Close()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%#v\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
