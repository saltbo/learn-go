package model

// goverter:converter
type Converter interface {
	Convert(source []Input) []Output
}

type XXX struct {
	Abc string
}

type Input struct {
	Name string
	Age  int
}
type Output struct {
	Name string
	Abc  string
	Age  int
}
