package mysql

import (
	"context"
	"database/sql"
	"fmt"
	mailing "mail_service/internal"

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
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCliente))
	query, args := courseSQLStruct.InsertInto(sqlClienteTable, sqlCliente{
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
