package sha

import (
	"crypto/sha256"
	"testing"
)

func TestCountDifferentBits(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"x", "X", 125},
		{"a", "A", 136},
	}
	for _, v := range tests {
		c1 := sha256.Sum256([]byte(v.a))
		c2 := sha256.Sum256([]byte(v.b))
		got := CountDifferentBits(c1, c2)
		if got != v.want {
			t.Errorf("CountDifferentBits(%v, %v) = %v, want = %v", v.a, v.b, got, v.want)
		}
	}
}
