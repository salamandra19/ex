package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "c", "d", "d", "e", "f"}
	fmt.Println(RemoveSame(s))
}

func RemoveSame(s []string) []string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
	}
	return s
}
