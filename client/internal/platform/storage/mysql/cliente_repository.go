package mysql

import (
	"context"
	"fmt"
	mailing "mail_service/internal"
	"mail_service/internal/kit/criteriamanager"

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

func (r *ClienteRepository) Update(ctx context.Context, cliente mailing.Cliente) error {
	sqlClienteModel := sqlCliente{
		Id:     cliente.ID().Value(),
		Nombre: cliente.NOMBRE().Value(),
		IdPlan: cliente.IDPLAN().Value(),
	}

	err := r.db.WithContext(ctx).Update(sqlClienteModel.Id, &sqlClienteModel).Error
	if err != nil {
		return fmt.Errorf("error trying to modify cliente on database: %v", err)
	}

	return nil
}

func (r *ClienteRepository) Delete(ctx context.Context, cliente mailing.Cliente) error {
	err := r.db.WithContext(ctx).Delete(&cliente).Error
	if err != nil {
		return fmt.Errorf("error trying to delete cliente on database: %v", err)
	}
	return err
}

func (r *ClienteRepository) Find(ctx context.Context, id string) (*mailing.Cliente, error) {
	var client mailing.Cliente
	err := r.db.WithContext(ctx).First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// Save implements the mooc.ClienteRepository interface.
func (r *ClienteRepository) Get(ctx context.Context, criteria criteriamanager.Criteria) (mailing.ClientesResponse, error) {

	var mysqlResponse []sqlClienteResponse

	err := selectFacade(
		ctx,
		r.db,
		criteria,
		sqlCliente{},
		&mysqlResponse,
		"clientes.*, planes.*",
		[]string{"INNER JOIN planes ON clientes.id_plan = planes.id"},
	)

	if err != nil {
		return nil, fmt.Errorf("error trying to get cliente on database: %v", err)
	}

	var responseData mailing.ClientesResponse

	for _, item := range mysqlResponse {
		responseData = append(responseData, func() mailing.ClienteResponse {
			var responseElem mailing.ClienteResponse
			convertStructToDomain(item, &responseElem)
			return responseElem
		}())
	}

	return responseData, nil
}
