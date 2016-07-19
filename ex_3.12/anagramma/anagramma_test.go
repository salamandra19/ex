package anagramma

import "testing"

func TestAnagramma(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"sort", "tors", true},
		{"sort", "tost", false},
		{"sotr", "sotr", true},
		{"sotr", "so", false},
		{"sotr", "sorted", false},
		{"", "", true},
		{"", "so", false},
		{"34", "43", true},
	}

	for _, v := range tests {
		got := Anagramma(v.a, v.b)
		if got != v.want {
			t.Errorf("Anagramma (%v, %v) = %v, want = %v", v.a, v.b, got, v.want)
		}
	}
}
