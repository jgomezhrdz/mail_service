package criteria

import (
	"fmt"
	"strconv"
	"strings"
)

// Filter represents the structure of the filter parameters
type Filter struct {
	Campo    string
	Operador string
	Valor    string
}

func ParseConditions(params [][]Filter) string {
	var conditions []string
	for _, elem := range params {
		var subConditions []string
		for _, filter := range elem {
			condition := generateCondition(&filter)
			if condition != "" {
				subConditions = append(subConditions, condition)
			}
		}
		// Combine conditions within the same subarray with OR
		subCondition := strings.Join(subConditions, " OR ")
		conditions = append(conditions, subCondition)
	}

	// Combine conditions between arrays with AND
	return strings.Join(conditions, " AND ")
}

// generateCondition generates a GORM condition string based on the filter parameters
func generateCondition(filter *Filter) string {
	switch filter.Operador {
	case "=":
		return fmt.Sprintf("%s = %s", filter.Campo, filter.Valor)
	case ">":
		val, err := strconv.Atoi(filter.Valor)
		if err == nil {
			return fmt.Sprintf("%s > %d", filter.Campo, val)
		}
	case "<":
		val, err := strconv.Atoi(filter.Valor)
		if err == nil {
			return fmt.Sprintf("%s < %d", filter.Campo, val)
		}
		// Add more cases for other operators as needed
	}
	return ""
}
