package sha

import (
	"crypto/sha256"
	"testing"
)

func TestSha(t *testing.T) {
	tests := []struct {
		a   string
		b   string
		sum int
	}{
		{"x", "X", 125},
		{"a", "A", 136},
	}
	for _, v := range tests {
		c1 := sha256.Sum256([]byte(v.a))
		c2 := sha256.Sum256([]byte(v.b))
		got := Sha(c1, c2)
		if got != v.sum {
			t.Errorf("Sha (%v, %v) = %v, sum = %v", v.a, v.b, got, v.sum)
		}
	}
}
