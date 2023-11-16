package criteria

import (
	"reflect"
	"strings"
)

func convertToSQLCliente(originalData interface{}, targetStruct interface{}) interface{} {
	// Get the types of the original and target structures
	originalType := reflect.TypeOf(originalData)
	sqlType := reflect.TypeOf(targetStruct)

	// Create a new instance of the target structure
	sqlClienteData := reflect.New(sqlType).Elem()

	// Iterate through the fields of the target structure
	for i := 0; i < sqlType.NumField(); i++ {
		// Get the field from the target structure
		targetField := sqlType.Field(i)

		// Extract the desired column name from the struct tag
		columnName := getColumnName(targetField)

		// Find the corresponding field in the original structure
		originalField, found := originalType.FieldByNameFunc(func(fieldName string) bool {
			return strings.EqualFold(fieldName, columnName)
		})

		// If the corresponding field is found, set its value in the target structure
		if found {
			value := reflect.ValueOf(originalData).FieldByName(originalField.Name).Interface()
			sqlClienteData.Field(i).Set(reflect.ValueOf(value))
		}
	}

	return sqlClienteData.Interface()
}

func getColumnName(field reflect.StructField) string {
	// Get the value of the "db" tag in the struct field tag
	tagValue := field.Tag.Get("db")

	// Split the tag value by "," and take the first part
	parts := strings.Split(tagValue, ",")
	return parts[0]
}
