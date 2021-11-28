package utils

import (
	"user_service/models"

	"github.com/go-playground/validator/v10"
)

func ValidatorStruct(i interface{}) []*models.ErrorResponse {
	var errors []*models.ErrorResponse

	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			var el models.ErrorResponse
			el.FailedField = v.StructNamespace()
			el.Tag = v.Tag()
			el.Value = v.Param()
			errors = append(errors, &el)

		}
	}

	return errors
}
