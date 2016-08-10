package removesame

import (
	"reflect"
	"testing"
)

func TestRemoveSame(t *testing.T) {
	tests := []struct {
		a    []string
		want []string
	}{
		{[]string{"aaa"}, []string{"aaa"}},
		{[]string{"a", "a", "a", "b"}, []string{"a", "b"}},
		{[]string{"bb", "bb"}, []string{"bb"}},
		{[]string{"bb", "aa", "bb", "aa"}, []string{"bb", "aa", "bb", "aa"}},
		{[]string{"bb", "bb", "aa", "aa"}, []string{"bb", "aa"}},
		{[]string{"aaa", "bbb", "ccc", "ccc", "ccc", "ddd", "ddd", "ddd", "ddd", "ccc", "bbb"}, []string{"aaa", "bbb", "ccc", "ddd", "ccc", "bbb"}},
		{[]string{}, []string{}},
	}
	for _, v := range tests {
		got := RemoveSame(v.a)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("RemoveSame %v = %v want %v", v.a, got, v.want)
		}
	}
}
