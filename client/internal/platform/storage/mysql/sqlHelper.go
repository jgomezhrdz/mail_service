package mysql

import (
	"context"
	"mail_service/internal/kit/criteriamanager"
	"reflect"

	"gorm.io/gorm"
)

type SQLEntity interface {
	convertSQLToDomain() (interface{}, error)
	TableName() string
}

func selectFacade(ctx context.Context, db *gorm.DB, criteria criteriamanager.Criteria, model interface{}, response interface{}, selectQuery string, joinQuery []string) error {
	query, values := criteriamanager.ParseConditions(criteria.GETFILTROS())

	queryBuilder := db.
		WithContext(ctx).
		Model(&model).
		Select(selectQuery)

	for _, item := range joinQuery {
		queryBuilder.Joins(item)
	}

	if limit := criteria.GETLIMIT(); limit != nil {
		queryBuilder.Limit(*limit)
	}

	if offset := criteria.GETOFFSET(); offset != nil {
		queryBuilder.Offset(*offset)
	}

	return queryBuilder.Where(query, values...).Scan(response).Error
}

func convertStructToDomain(src interface{}, dest interface{}) {
	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.Field(i)
		// Check if the field implements the SQLEntity interface
		if sqlEntity, ok := srcField.Interface().(SQLEntity); ok {
			// Call the convertSQLToDomain method and set the value in the destination field
			result, _ := sqlEntity.convertSQLToDomain()
			destField.Set(reflect.ValueOf(result))
		}
	}
}
