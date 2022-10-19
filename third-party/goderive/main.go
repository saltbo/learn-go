//go:generate go run github.com/awalterschulze/goderive .

package main

import "fmt"

type MyStruct struct {
	Int64     int64
	StringPtr *string
}

func (this *MyStruct) Equal(that *MyStruct) bool {
	return deriveEqual(this, that)
}

func main() {
	a := 10
	b := 20
	fmt.Println(deriveMax(a, b))
}
