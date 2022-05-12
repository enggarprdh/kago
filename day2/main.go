package main

import (
	"flag"
	"fmt"
	"kago/gt"
)

var (
	txt = flag.String("txt", "", "-txt=masukkan text yang akan di hitung")
)

func main() {
	flag.Parse()
	fmt.Println("Kata yang dimasukkan : ", *txt)
	v, c := gt.GetCharCount(*txt)
	fmt.Println("vocals:", len(v))
	fmt.Println("consonans:", len(c))
}
