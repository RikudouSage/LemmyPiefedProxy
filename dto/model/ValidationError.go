package model

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	Err   validator.ValidationErrors
	Input any
}

func (receiver *ValidationError) Error() string {
	return receiver.Err.Error()
}

func NewValidationError() *ValidationError {
	return &ValidationError{}
}
