package mysql

import mailing "mail_service/internal"

type sqlClienteResponse struct {
	Client sqlCliente `gorm:"embedded"`
	Plan   sqlPlan    `gorm:"embedded"`
}

type sqlCliente struct {
	Id     string `db:"id"`
	Nombre string `db:"nombre"`
	IdPlan string `db:"id_plan"`
}

func (sqlCliente) TableName() string {
	return "clientes"
}

func (sqlCliente sqlCliente) convertSQLToDomain() (interface{}, error) {
	return mailing.NewCliente(sqlCliente.Id, sqlCliente.Nombre, sqlCliente.IdPlan)
}
