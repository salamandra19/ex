package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
}

func comma(s string) string {
	var b bytes.Buffer
	x := len(s) % 3
	if x == 0 {
		x = 3
	}
	for i := x; len(s) > 0; {
		if len(s) > 3 {
			fmt.Fprintf(&b, "%s", s[:i])
			b.WriteString(",")
			s = s[i:]
			i = 3
		} else {
			b.WriteString((s[:i]))
			break
		}
	}
	return b.String()
}
