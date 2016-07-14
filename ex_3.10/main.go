package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma2(s))
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

func comma1(s string) string {
	var b bytes.Buffer
	for i := len(s) % 3; len(s) > 0; i = 3 {
		if b.Len() > 0 {
			b.WriteString(",")
		}
		b.WriteString(s[:i])
		s = s[i:]
	}
	return b.String()
}

func comma2(s string) string {
	var b bytes.Buffer
	var n = 0
	for i := len(s) % 3; len(s) >= i; i += 3 {
		if b.Len() > 0 {
			b.WriteString(",")
		}
		b.WriteString(s[n:i])
		n = i
	}
	return b.String()
}
