package mailing

import (
	"mail_service/internal/kit/event"
)

const ClienteCreatedEventType event.Type = "events.cliente.created"

type ClienteCreatedEvent struct {
	event.BaseEvent

	id     string
	nombre string
	idPlan string
}

func NewClienteCreatedEvent(id, nombre, idPlan string) ClienteCreatedEvent {
	return ClienteCreatedEvent{
		id:     id,
		nombre: nombre,
		idPlan: idPlan,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e ClienteCreatedEvent) Type() event.Type {
	return ClienteCreatedEventType
}

func (e ClienteCreatedEvent) CourseID() string {
	return e.id
}

func (e ClienteCreatedEvent) CourseName() string {
	return e.nombre
}

func (e ClienteCreatedEvent) CourseDuration() string {
	return e.idPlan
}
