package main

import (
	"fmt"
	"kago/gt"
)

var (
	f32 float32
	f64 float64
)

func main() {
	i := gt.IntData(100)
	s := gt.StringData(5)
	b := true
	f32 = 80
	f64 = 1200

	fmt.Println(i, s, b, f32, f64)
	fmt.Printf("%T %T %T %T %T\n", i, s, b, f32, f64)

	fmt.Printf("%s = %d adalah %T , %s = %s adalah %T", "i", i, i, "s", s, s)
}
