// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import model "goverter/model"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source []model.Input) []model.Output {
	modelOutputList := make([]model.Output, len(source))
	for i := 0; i < len(source); i++ {
		modelOutputList[i] = c.modelInputToModelOutput(source[i])
	}
	return modelOutputList
}
func (c *ConverterImpl) modelInputToModelOutput(source model.Input) model.Output {
	var modelOutput model.Output
	modelOutput.Name = source.Name
	modelOutput.Age = source.Age
	return modelOutput
}