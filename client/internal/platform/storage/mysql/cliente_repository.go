package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	clientSQLStruct := sqlbuilder.NewStruct(new(sqlCliente))
	query, args := clientSQLStruct.InsertInto(sqlClienteTable, sqlCliente{
		Id:     cliente.ID().Value(),
		Nombre: cliente.NOMBRE().Value(),
		IdPlan: cliente.IDPLAN().Value(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist cliente on database: %v", err)
	}

	return nil
}

// Save implements the mooc.ClienteRepository interface.
func (r *ClienteRepository) Get(ctx context.Context) ([]struct {
	Client mailing.Cliente
	Plan   mailing.Plan
}, error) {
	sb :=
		sqlbuilder.Select("clientes.*", "planes.*").From("clientes").JoinWithOption(sqlbuilder.InnerJoin, "planes", "clientes.id_plan = planes.id")
	sql, args := sb.Build()

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error trying to get cliente on database: %v", err)
	}
	defer rows.Close()

	var responseData []struct {
		Client mailing.Cliente
		Plan   mailing.Plan
	}

	for rows.Next() {
		var cliente sqlCliente
		var plan sqlPlan

		// Assuming the fields of Cliente and Plan are of types that Scan understands (e.g., int, string, etc.)
		err := rows.Scan(&cliente.Id, &cliente.Nombre, &cliente.IdPlan, &plan.Id, &plan.Nombre, &plan.QuotaDay, &plan.QuotaMonth)
		if err != nil {
			log.Fatal(err)
		}

		mailingCliente, err := convertSQLClienteToMailingCliente(cliente)
		if err != nil {
			return nil, fmt.Errorf("error converting sqlCliente to mailing.Cliente: %v", cliente)
		}

		mailingPlan, err := convertSQLPlanToMailingPlan(plan)
		if err != nil {
			return nil, fmt.Errorf("error converting sqlPlan to mailing.Plan: %v", err)
		}

		responseData = append(responseData, struct {
			Client mailing.Cliente
			Plan   mailing.Plan
		}{Client: mailingCliente, Plan: mailingPlan})

		fmt.Printf("Customer: %+v, Planes: %+v", cliente, plan)
	}

	return responseData, nil
}
