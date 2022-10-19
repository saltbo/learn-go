package main

import (
	"goverter/generated"
	"goverter/model"
)

func main() {
	in := model.Input{
		Name: "",
		Age:  0,
	}

	var c model.Converter
	c = &generated.ConverterImpl{}
	c.Convert([]model.Input{in})
}
