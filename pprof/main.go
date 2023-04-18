package main

import (
	"fmt"
	"runtime"
)

type T0 struct {
	Abc string `json:"abc"`
}

type T1 struct {
	Name string `json:"name"`

	Meta T0 `json:"meta"`
}

type T2 struct {
	TT *T0 `json:"tt"`
}

func main() {
	meta := T1{
		Name: "aaa",
		Meta: T0{
			Abc: "111",
		},
	}
	println(&meta.Meta)
	meta = T1{}

	abc := T2{TT: &meta.Meta}
	println(abc.TT)
	fmt.Print(abc.TT.Abc)
	runtime.GC()

}
