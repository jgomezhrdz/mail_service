package types

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UUID struct {
	value string
}

var ErrInvalidID = errors.New("the field ID is not a valid UUID")

func NewUUID(value string) (UUID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UUID{}, fmt.Errorf("%w: %s", ErrInvalidID, value)
	}

	return UUID{
		value: v.String(),
	}, nil
}

func (id UUID) Value() string {
	return id.value
}
