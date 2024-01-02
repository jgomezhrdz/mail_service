package mysql

import (
	"reflect"
)

type SQLEntity interface {
	convertSQLToDomain() (interface{}, error)
	TableName() string
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
