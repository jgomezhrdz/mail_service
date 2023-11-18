package cliente_services

import (
	"context"
	mailing "mail_service/internal"
	"mail_service/internal/kit/event"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type ClienteService struct {
	clienteRepository mailing.ClienteRepository
	eventBus          event.Bus
}

// NewCourseService returns the default Service interface implementation.
func NewClienteService(clienteRepository mailing.ClienteRepository, eventBus event.Bus) ClienteService {
	return ClienteService{
		clienteRepository: clienteRepository,
		eventBus:          eventBus,
	}
}

func (s ClienteService) CreateCliente(ctx context.Context, id, nombre, duration string) error {
	cliente, err := mailing.NewCliente(id, nombre, duration)
	if err != nil {
		return err
	}

	if err := s.clienteRepository.Save(ctx, cliente); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, cliente.PullEvents())
}

func (s ClienteService) GetCliente(ctx context.Context) (mailing.ClientesResponse, error) {
	result, err := s.clienteRepository.Get(ctx)

	var jsonData mailing.ClientesResponse

	for _, instance := range result {
		clientJSON := instance.Client.TOJSON() // Assuming ToJSON is a method in mailing.Cliente
		planJSON := instance.Plan.TOJSON()     // Assuming ToJSON is a method in mailing.Plan

		jsonData = append(jsonData, mailing.ClienteResponse{
			Client: clientJSON, Plan: planJSON,
		})
	}

	return jsonData, err
}
