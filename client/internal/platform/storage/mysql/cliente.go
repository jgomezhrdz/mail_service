package mysql

const (
	sqlClienteTable = "clientes"
)

type sqlCliente struct {
	Id     string `db:"id"`
	Name   string `db:"nombre"`
	IdPlan string `db:"id_plan"`
}
