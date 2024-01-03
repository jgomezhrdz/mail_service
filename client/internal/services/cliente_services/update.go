package cliente_services

import (
	"context"
)

func (s ClienteService) UpdateCliente(ctx context.Context, id, nombre, idPlan string) error {
	cliente, err := s.clienteRepository.Find(ctx, id)
	if err != nil {
		return err
	}

	err = cliente.UPDATE(nombre, idPlan)
	if err != nil {
		return err
	}

	if err := s.clienteRepository.Update(ctx, cliente); err != nil {
		return err
	}

	return nil
}
