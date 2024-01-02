package mysql

import (
	"context"
	"fmt"
	mailing "mail_service/internal"
	"mail_service/internal/kit/criteria"

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
func (r *ClienteRepository) Get(ctx context.Context, filters [][]criteria.Filter) (mailing.ClientesResponse, error) {

	var mysqlResponse []sqlClienteResponse

	query, values := criteria.ParseConditions(filters)

	err := r.db.
		WithContext(ctx).
		Model(&sqlCliente{}).
		Select("clientes.*, planes.*").
		Joins("INNER JOIN planes ON clientes.id_plan = planes.id").Where(query, values...).
		Scan(&mysqlResponse).Error

	if err != nil {
		return nil, fmt.Errorf("error trying to get cliente on database: %v", err)
	}

	var responseData mailing.ClientesResponse

	for _, item := range mysqlResponse {
		var responseElem mailing.ClienteResponse
		convertStructToDomain(item, &responseElem)
		responseData = append(responseData, responseElem)
	}
	return responseData, nil
}
