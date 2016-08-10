package removesame

func RemoveSame(s []string) []string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
	}
	return s
}
