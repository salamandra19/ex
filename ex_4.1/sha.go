package sha

var pc [256]byte

// pc[i] Делением i/2 получаем результат суммы выставленных битов предыдущего
// по величине числа.
// Прибавляем (если выставлен) бит текущего.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Sha(x, y [32]byte) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += int(pc[x[i]^y[i]])
	}
	return sum
}
