package cliente_services

import (
	"context"
	mailing "mail_service/internal"
)

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
