package cliente_services

import (
	"context"
	mailing "mail_service/internal"
)

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
