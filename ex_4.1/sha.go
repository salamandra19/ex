package sha

var pc [256]byte

// pc[i] Делением i/2 получаем результат суммы выставленных битов вдвое
// меньшего числа, т.е сумму всех битов текущего числа, кроме самого
// младшего (правого) бита.
// Прибавляем младший (правый) бит текущего числа (если выставлен).
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CountDifferentBits(x, y [32]byte) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += int(pc[x[i]^y[i]])
	}
	return sum
}
