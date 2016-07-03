package popcount

import "testing"

func TestPopCount(t *testing.T) {
	cases := []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{5, 2},
		{12, 2},
		{0xFFF000, 12},
		{^uint64(0), 64},
	}
	for _, v := range cases {
		got := PopCount(v.x)
		if got != v.want {
			t.Errorf("PopCount(%d) = %d, want %d", v.x, got, v.want)
		}
	}
}
