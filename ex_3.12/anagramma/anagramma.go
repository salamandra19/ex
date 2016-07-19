package anagramma

import (
	"sort"
	"strings"
)

func Anagramma(a, b string) bool {
	as := strings.Split(a, "")
	bs := strings.Split(b, "")
	sort.Strings(as)
	sort.Strings(bs)
	ass := strings.Join(as, "")
	bss := strings.Join(bs, "")
	if ass == bss {
		return true
	}
	return false
}
