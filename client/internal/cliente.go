package mailing

import (
	"context"
	types "mail_service/internal/kit/types"
)

type Cliente struct {
	id     types.UUID
	nombre types.NonEmptyString
	idPlan types.UUID
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ClienteRepository
type ClienteRepository interface {
	Get(ctx context.Context) ([]struct {
		Client Cliente
		Plan   Plan
	}, error)
	Save(ctx context.Context, cliente Cliente) error
}

type ClientesResponse []ClienteResponse
type ClienteResponse struct {
	Client map[string]interface{}
	Plan   map[string]interface{}
}

// NewCourse creates a new course.
func NewCliente(id string, nombre string, idPlan string) (Cliente, error) {

	idVO, err := types.NewUUID(id)
	if err != nil {
		return Cliente{}, err
	}

	nombreVO, err := types.NewNonEmptyString(nombre)
	if err != nil {
		return Cliente{}, err
	}

	planVO, err := types.NewUUID(idPlan)
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
func (c Cliente) ID() types.UUID {
	return c.id
}

func (c Cliente) NOMBRE() types.NonEmptyString {
	return c.nombre
}

func (c Cliente) IDPLAN() types.UUID {
	return c.idPlan
}

func (c Cliente) TOJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":     c.ID().Value(),
		"nombre": c.NOMBRE().Value(),
		"idPlan": c.IDPLAN().Value(),
		// Add other fields as needed
	}
}
