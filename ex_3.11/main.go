package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(detail(s))
	}
}

func detail(s string) string {
	var rest, sign string
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		sign, s = string(s[0]), s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s, rest = s[:i], s[i:]
			break
		}
	}
	return sign + comma(s) + rest
}

func comma(s string) string {
	var b bytes.Buffer
	for start, end := 0, len(s)%3; end <= len(s); start, end = end, end+3 {
		if b.Len() > 0 {
			b.WriteString(",")
		}
		b.WriteString(s[start:end])
	}
	return b.String()
}
