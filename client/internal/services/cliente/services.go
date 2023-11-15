package cliente_services

import (
	"context"
	mailing "mail_service/internal"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type ClienteService struct {
	clienteRepository mailing.ClienteRepository
}

// NewCourseService returns the default Service interface implementation.
func NewClienteService(clienteRepository mailing.ClienteRepository) ClienteService {
	return ClienteService{
		clienteRepository: clienteRepository,
	}
}

func (s ClienteService) CreateCliente(ctx context.Context, id, nombre, duration string) error {
	course, err := mailing.NewCliente(id, nombre, duration)
	if err != nil {
		return err
	}
	return s.clienteRepository.Save(ctx, course)
}

func (s ClienteService) GetCliente(ctx context.Context) (mailing.ClientesResponse, error) {
	result, err := s.clienteRepository.Get(ctx)

	var jsonData mailing.ClientesResponse

	for _, instance := range result {
		clientJSON := mailing.Cliente.TOJSON(instance.Client) // Assuming ToJSON is a method in mailing.Cliente
		planJSON := mailing.Plan.TOJSON(instance.Plan)        // Assuming ToJSON is a method in mailing.Plan

		jsonData = append(jsonData, mailing.ClienteResponse{
			Client: clientJSON, Plan: planJSON,
		})
	}

	return jsonData, err
}
