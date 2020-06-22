package main

import (
	"fmt"
	"github.com/go-basic/fixedlist"
)

func main() {
	f := fixedlist.NewFixedList(2)
	f.Add("a")
	f.Add("b")
	fmt.Println(f.Data())
	f.Add("c")
	fmt.Println(f.Data())
	f.Add("d")
	fmt.Println(f.Data())
	f.Add("e")
	fmt.Println(f.Data())
}

/**
[a b]
[b c]
[c d]
[d e]
 */