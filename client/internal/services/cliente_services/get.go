package cliente_services

import (
	"context"
	mailing "mail_service/internal"
	"mail_service/internal/kit/criteria"
)

func (s ClienteService) GetCliente(ctx context.Context) (mailing.ClientesResponse, error) {
	result, err := s.clienteRepository.Get(ctx, [][]criteria.Filter{})

	var jsonData mailing.ClientesResponse

	if len(result) == 0 {
		jsonData = []mailing.ClienteResponse{}
	} else {
		jsonData = result
	}

	return jsonData, err
}
