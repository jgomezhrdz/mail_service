package tools

import (
	"fmt"
	"reflect"
)

func StructKeys(s interface{}) []interface{} {
	var keys []interface{}

	val := reflect.ValueOf(s)

	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return keys
	}

	for i := 0; i < val.NumField(); i++ {
		keys = append(keys, val.Field(i).Addr().Interface())
	}

	return keys
}

func StructColumnMap(s interface{}) []interface{} {
	columnMap := make(map[string]interface{})

	val := reflect.ValueOf(s).Elem() // Get the value of the struct (assuming s is a pointer to the struct)

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		columnMap[fieldName] = val.Field(i).Addr().Interface()
	}

	dest := make([]interface{}, len(columnMap))
	i := 0
	for _, v := range columnMap {
		dest[i] = v
		i++
	}
	return dest
}
