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
	if x != 0 {
		fmt.Fprintf(&b, "%s", s[:x])
	}
	s = s[x:]
	for len(s) > 0 {
		if b.Len() > 0 {
			b.WriteString(",")
		}
		fmt.Fprintf(&b, "%s", s[:3])
		s = s[3:]
	}
	return b.String()
}
