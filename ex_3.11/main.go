package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var rest, sign string
	for _, s := range os.Args[1:] {
		rest, sign = "", ""
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '.' {
				rest = s[i:]
				s = s[:i]
			}
			if s[0] == '-' {
				sign = "-"
				s = s[1:]
			}
		}
		fmt.Println(sign + comma(s) + rest)
	}
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

/*
func comma(s string) string {
	var b bytes.Buffer
	for i := len(s) % 3; len(s) > 3; i = 3 {
		if i == 0 {
			i = 3
		}
		ss, _ := strconv.ParseFloat(s[:i], 64)
		fmt.Fprintf(&b, "%v", ss)
		b.WriteString(",")
		s = s[i:]
	}
	b.WriteString(s)
	return b.String()
}

	ss, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}

	n := len(ss)
	if n <= 3 {
		return ss
	}
	return comma(ss[:n-3]) + "," + ss[n-3:]
*/
