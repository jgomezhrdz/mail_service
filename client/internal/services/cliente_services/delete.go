package cliente_services

import (
	"context"
)

func (s ClienteService) DeleteCliente(ctx context.Context, id string) error {

	if err := s.clienteRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
