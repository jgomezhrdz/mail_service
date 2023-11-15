package mysql

import (
	"context"
	"fmt"
	mailing "mail_service/internal"

	"gorm.io/gorm"
)

// ClienteRepository is a MySQL mooc.ClienteRepository implementation.
type ClienteRepository struct {
	db *gorm.DB
}

// NewClienteRepository initializes a MySQL-based implementation of mooc.ClienteRepository.
func NewClienteRepository(db *gorm.DB) *ClienteRepository {
	return &ClienteRepository{
		db: db,
	}
}

// Save implements the mooc.ClienteRepository interface.
func (r *ClienteRepository) Save(ctx context.Context, cliente mailing.Cliente) error {
	sqlClienteModel := sqlCliente{
		Id:     cliente.ID().Value(),
		Nombre: cliente.NOMBRE().Value(),
		IdPlan: cliente.IDPLAN().Value(),
	}

	err := r.db.WithContext(ctx).Create(&sqlClienteModel).Error
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

	var mysqlResponse []struct {
		Client sqlCliente `gorm:"embedded"`
		Plan   sqlPlan    `gorm:"embedded"`
	}

	err := r.db.
		WithContext(ctx).
		Model(&mailing.Cliente{}).
		Select("clientes.*, planes.*").
		Joins("INNER JOIN planes ON clientes.id_plan = planes.id").
		Scan(&mysqlResponse).Error

	if err != nil {
		return nil, fmt.Errorf("error trying to get cliente on database: %v", err)
	}

	var responseData []struct {
		Client mailing.Cliente
		Plan   mailing.Plan
	}

	for _, item := range mysqlResponse {
		mailingCliente, err := convertSQLClienteToMailingCliente(item.Client)
		if err != nil {
			return nil, fmt.Errorf("error converting sqlCliente to mailing.Cliente: %v", item.Client)
		}

		mailingPlan, err := convertSQLPlanToMailingPlan(item.Plan)
		if err != nil {
			return nil, fmt.Errorf("error converting sqlPlan to mailing.Plan: %v", err)
		}

		responseData = append(responseData, struct {
			Client mailing.Cliente
			Plan   mailing.Plan
		}{Client: mailingCliente, Plan: mailingPlan})
	}

	return responseData, nil
}
