package anagramma

import "testing"

func TestAnagramma(t *testing.T) {
	a := []string{"s", "o", "r", "t"}
	b := []string{"t", "o", "r", "s"}
	if !Anagramma(a, b) {
		t.Error("Anagramma %v %v =false", a, b)
	}
}
