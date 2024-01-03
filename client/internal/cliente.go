package mailing

import (
	"context"
	"mail_service/internal/kit/criteriamanager"
	"mail_service/internal/kit/event"
	types "mail_service/internal/kit/types"
)

type Cliente struct {
	id     types.UUID
	nombre types.NonEmptyString
	idPlan types.UUID

	events []event.Event
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ClienteRepository
type ClienteRepository interface {
	Get(ctx context.Context, criteria criteriamanager.Criteria) (ClientesResponse, error)
	Find(ctx context.Context, id string) (Cliente, error)
	Save(ctx context.Context, cliente Cliente) error
	Update(ctx context.Context, cliente Cliente) error
	Delete(ctx context.Context, id string) error
}

type ClientesResponse []ClienteResponse
type ClienteResponse struct {
	Client Cliente
	Plan   Plan
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

func (c Cliente) UPDATE(nombre string, idPlan string) (Cliente, error) {
	nombreVO, err := types.NewNonEmptyString(nombre)
	if err != nil {
		return c, err
	}
	planVO, err := types.NewUUID(idPlan)
	if err != nil {
		return c, err
	}
	c.nombre = nombreVO
	c.idPlan = planVO

	return c, nil
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
	}
}

// Record records a new domain event.
func (c *Cliente) Record(evt event.Event) {
	c.events = append(c.events, evt)
}

// PullEvents returns all the recorded domain events.
func (c Cliente) PullEvents() []event.Event {
	evt := c.events
	c.events = []event.Event{}

	return evt
}
