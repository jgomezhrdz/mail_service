package cliente_services

import (
	"context"
	mailing "mail_service/internal"
	criteriamanager "mail_service/internal/kit/criteriamanager"
)

func (s ClienteService) GetCliente(ctx context.Context, queryParam map[string][]string) (mailing.ClientesResponse, error) {

	result, err := s.clienteRepository.Get(ctx, criteriamanager.CriteriaFromRequest(queryParam))

	var jsonData mailing.ClientesResponse

	if len(result) == 0 || result == nil {
		jsonData = []mailing.ClienteResponse{}
	} else {
		jsonData = result
	}

	return jsonData, err
}
