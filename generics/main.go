package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int | int64 | float64
}

func Sum[K comparable, V Number](m map[K]V) V {

	var s V
	for _, v := range m {
		s += v
	}

	return s
}

type ABC struct {
	Name string
}

type BCD struct {
	Age int
}

type AA interface {
	ABC | BCD
}

// type A[T AA] struct {
// }
//
// func (receiver A[T]) aaa(m T) {
// 	switch k := m.(type) {
// 	case ABC:
// 		fmt.Println("ABC")
// 	case BCD:
// 		fmt.Println("BCD")
// 	}
// 	fmt.Println(m < 10)
// }

type XX struct {
	Name string
}

type X1 struct {
	Age string
}

type BB interface {
	XX | X1
}

func Test[T AA, T2 BB](a T) T2 {
	x := T2{}

	return T2[BB]{Name: "xx"}
}

func main() {
	Test[ABC, XX](ABC{})

	// a := A[BCD]{}
	// a.aaa(BCD{})

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		Sum(ints),
		Sum(floats),
		Sum(map[int64]int{
			1000000: 2,
		}),
	)
}
