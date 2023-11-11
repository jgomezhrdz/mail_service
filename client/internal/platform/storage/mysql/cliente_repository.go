package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	mailing "mail_service/internal"
	tools "mail_service/internal/platform/shared/tools"

	"github.com/huandu/go-sqlbuilder"
)

// ClienteRepository is a MySQL mooc.ClienteRepository implementation.
type ClienteRepository struct {
	db *sql.DB
}

// NewClienteRepository initializes a MySQL-based implementation of mooc.ClienteRepository.
func NewClienteRepository(db *sql.DB) *ClienteRepository {
	return &ClienteRepository{
		db: db,
	}
}

// Save implements the mooc.ClienteRepository interface.
func (r *ClienteRepository) Save(ctx context.Context, cliente mailing.Cliente) error {
	clientSQLStruct := sqlbuilder.NewStruct(new(sqlCliente))
	query, args := clientSQLStruct.InsertInto(sqlClienteTable, sqlCliente{
		Id:     cliente.ID().String(),
		Name:   cliente.NOMBRE().String(),
		IdPlan: cliente.IDPLAN().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist cliente on database: %v", err)
	}

	return nil
}

// Save implements the mooc.ClienteRepository interface.
func (r *ClienteRepository) Get(ctx context.Context) error {
	sb :=
		sqlbuilder.Select("clientes.*", "planes.*").From("clientes").JoinWithOption(sqlbuilder.InnerJoin, "planes", "clientes.id_plan = planes.id")
	sql, args := sb.Build()

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("error trying to get cliente on database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cliente sqlCliente
		var plan sqlPlan

		clienteMap := tools.StructColumnMap(&cliente)
		planMap := tools.StructColumnMap(&plan)

		// Assuming the fields of Cliente and Plan are of types that Scan understands (e.g., int, string, etc.)
		err := rows.Scan(append(clienteMap, planMap...)...)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Customer: %+v, Planes: %+v", cliente, plan)
	}

	return nil
}
