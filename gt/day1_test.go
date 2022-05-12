package gt_test

import (
	"kago/gt"
	"testing"
)

func TestInt(t *testing.T) {
	max := 20
	i := gt.IntData(max)
	if i > max {
		t.Fatal("salah > 10")
	}

}

func TestString(t *testing.T) {
	max := 20
	s := gt.StringData(max)
	if len(s) != max {
		t.Fatal("panjang tidak sama")
	}
}
