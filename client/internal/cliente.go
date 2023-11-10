package mailing

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")

// Course is the data structure that represents a course.
type Cliente struct {
	id     ClienteID
	nombre ClienteNombre
	idPlan PlanID
}

type ClienteID struct {
	value string
}

var ErrInvalidID = errors.New("the field ID is not a valid UUID")

func NewClienteID(value string) (ClienteID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ClienteID{}, fmt.Errorf("%w: %s", ErrInvalidID, value)
	}

	return ClienteID{
		value: v.String(),
	}, nil
}

func (id ClienteID) String() string {
	return id.value
}

type ClienteNombre struct {
	value string
}

var ErrEmptyNombre = errors.New("the field Nombre can not be empty")

func NewClienteNombre(value string) (ClienteNombre, error) {
	if value == "" {
		return ClienteNombre{}, ErrEmptyNombre
	}

	return ClienteNombre{
		value: value,
	}, nil
}

func (id ClienteNombre) String() string {
	return id.value
}

// PlanID represents the id plan related to the client.
type PlanID struct {
	value string
}

var ErrInvalidPlan = errors.New("the field Plan ID is not a valid UUID")

func NewClientePlan(value string) (PlanID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return PlanID{}, fmt.Errorf("%w: %s", ErrInvalidPlan, value)
	}

	return PlanID{
		value: v.String(),
	}, nil
}

func (id PlanID) String() string {
	return id.value
}

// NewCourse creates a new course.
func NewCliente(id string, nombre string, idPlan string) (Cliente, error) {

	idVO, err := NewClienteID(id)
	if err != nil {
		return Cliente{}, err
	}

	nombreVO, err := NewClienteNombre(nombre)
	if err != nil {
		return Cliente{}, err
	}

	planVO, err := NewClientePlan(idPlan)
	if err != nil {
		return Cliente{}, err
	}

	return Cliente{
		id:     idVO,
		nombre: nombreVO,
		idPlan: planVO,
	}, nil
}

// ID returns the course unique identifier.
func (c Cliente) ID() ClienteID {
	return c.id
}

func (c Cliente) NOMBRE() ClienteNombre {
	return c.nombre
}

func (c Cliente) IDPLAN() PlanID {
	return c.idPlan
}
