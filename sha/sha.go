package sha

func Sha(a, b [32]uint8) int {
	var s int
	for i := 0; i < 32; i++ {
		s = int(a[i] ^ b[i])
	}
	return s
}
