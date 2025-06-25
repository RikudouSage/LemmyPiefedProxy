package helper

import (
	"LemmyPiefedApi/dto/model"
	"LemmyPiefedApi/http"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ParseRequest[T any](request *http.Request) (*T, error) {
	var result T
	err := json.Unmarshal(request.Body, &result)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(result)
	if err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			return nil, &model.ValidationError{
				Err:   validationErrs,
				Input: result,
			}
		}
		return nil, err
	}

	return &result, nil
}
