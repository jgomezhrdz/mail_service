package mysql

import mailing "mail_service/internal"

const (
	sqlClienteTable = "clientes"
)

type sqlCliente struct {
	Id     string `db:"id"`
	Nombre string `db:"nombre"`
	IdPlan string `db:"id_plan"`
}

func convertSQLClienteToMailingCliente(sqlCliente sqlCliente) (mailing.Cliente, error) {
	cliente, err := mailing.NewCliente(sqlCliente.Id, sqlCliente.Nombre, sqlCliente.IdPlan)
	if err != nil {
		return cliente, err
	}
	return cliente, nil
}
