package sha

import (
	"crypto/sha256"
	"testing"
)

func testSha(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"x", "X", 125},
		{"a", "A", 136},
	}
	for _, v := range tests {
		as := sha256.Sum256([]byte(v.a))
		bs := sha256.Sum256([]byte(v.b))
		got := Sha(as, bs)
		if got != v.want {
			t.Errorf("Sha(%v,%v) = %v, want = %v", v.a, v.b, got, v.want)
		}
	}
}
