package types

import "errors"

type UnsignedInt struct {
	value int
}

var ErrNegativaString = errors.New("the field Nombre can not be negativa")

func NewUnsignedInt(value int) (UnsignedInt, error) {
	if value < 0 {
		return UnsignedInt{}, ErrNegativaString
	}

	return UnsignedInt{
		value: value,
	}, nil
}

func (id UnsignedInt) Value() int {
	return id.value
}
