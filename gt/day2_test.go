package gt_test

import (
	"fmt"
	"kago/gt"
	"testing"
)

func TestCharCount(t *testing.T) {
	word := "desin"
	v, c := gt.GetCharCount(word)

	fmt.Println(len(v))
	fmt.Println(len(c))
}
