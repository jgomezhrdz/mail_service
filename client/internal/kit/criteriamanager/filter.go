package criteriamanager

import (
	"fmt"
	"strings"
)

// Filter represents the structure of the filter parameters
type Filter struct {
	Campo    string
	Operador string
	Valor    string
}

func NewFilter(campo string, operador string, valor string) Filter {
	return Filter{Campo: campo, Operador: operador, Valor: valor}
}

func ParseConditions(params [][]Filter) (string, []interface{}) {
	var conditions []string
	var queryValues []interface{}
	for _, elem := range params {
		var subConditions []string
		for _, filter := range elem {
			condition, valor := generateCondition(&filter)
			queryValues = append(queryValues, valor)
			if condition != "" {
				subConditions = append(subConditions, condition)
			}
		}
		// Combine conditions within the same subarray with OR
		subCondition := strings.Join(subConditions, " OR ")
		conditions = append(conditions, subCondition)
	}

	// Combine conditions between arrays with AND
	return strings.Join(conditions, " AND "), queryValues
}

// generateCondition generates a GORM condition string based on the filter parameters
func generateCondition(filter *Filter) (string, string) {
	query := ""
	valor := filter.Valor
	//parse json values
	result := convertToJSONPath(filter.Campo)

	switch filter.Operador {
	case "=":
		query = fmt.Sprintf("%s = ?", result)
	case ">":
		query = fmt.Sprintf("%s > ?", result)
	case ">=":
		query = fmt.Sprintf("%s > ?", result)
	case "<":
		query = fmt.Sprintf("%s < ?", result)
	case "<=":
		query = fmt.Sprintf("%s <= ?", result)
	case "IS", "is":
		query = fmt.Sprintf("%s IS NULL", result)
	case "LIKE", "like":
		valor = "%" + valor + "%"
		query = fmt.Sprintf("%s LIKE ?", result)
	}

	return query, valor
}

func convertToJSONPath(input string) string {
	parts := strings.Split(input, "->")
	var expression string
	if len(parts) > 1 {
		expression = "JSON_EXTRACT("
		for i, part := range parts {
			if i > 1 {
				expression += fmt.Sprintf("JSON_EXTRACT(%s, '$.%s')", parts[i-1], part)
			}
		}
		expression += strings.Repeat(")", len(parts)-1)
	} else {
		expression = parts[0]
	}

	return expression
}
