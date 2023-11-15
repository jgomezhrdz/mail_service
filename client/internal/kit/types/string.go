package types

import "errors"

type NonEmptyString struct {
	value string
}

var ErrEmptyString = errors.New("the field Nombre can not be empty")

func NewNonEmptyString(value string) (NonEmptyString, error) {
	if value == "" {
		return NonEmptyString{}, ErrEmptyString
	}

	return NonEmptyString{
		value: value,
	}, nil
}

func (id NonEmptyString) Value() string {
	return id.value
}
