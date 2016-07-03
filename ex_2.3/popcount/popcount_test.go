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

		got = PopCountLoop(v.x)
		if got != v.want {
			t.Errorf("PopCountLoop(%d) = %d, want %d", v.x, got, v.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1122334455667788)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1122334455667788)
	}
}
