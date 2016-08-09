package main

import "fmt"

const (
	// _  = math.Pow10(3)
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB, float64(ZB), float64(YB))
}
