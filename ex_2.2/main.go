package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/salamandra19/ex/ex_2.1/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convert(input.Text())
		}
		err := input.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func convert(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("cf: %v", err)
		return
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
