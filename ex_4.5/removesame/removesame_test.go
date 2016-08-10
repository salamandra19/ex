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
		{[]string{"b"}, []string{"b"}},
		{[]string{}, []string{}},
	}
	for _, v := range tests {
		got := RemoveSame(v.a)
		if reflect.DeepEqual(got, v.want) == false {
			t.Errorf("RemoveSame %v = %v want %v", v.a, got, v.want)
		}
	}
}
