/*
# This validation package is mainly for custom validation

  - See `custom.go` for custom validation tag
  - See `validate.go` for custom validation message
*/
package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateIsType(fl validator.FieldLevel) bool {
	return false
}
