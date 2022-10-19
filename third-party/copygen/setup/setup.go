/* Specify the name of the generated file's package. */
package setup

import (
	"copygen/domain"
	"copygen/model"
)

/* Copygen defines the functions that will be generated. */
type Copygen interface {
	// custom see table below for options
	ModelsToDomain(*model.Account, *model.User) *domain.Account
}
