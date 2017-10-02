package sha

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Sha(a, b [32]byte) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += int(pc[(a[i] ^ b[i])])
	}
	return sum
}
