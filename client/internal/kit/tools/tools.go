package tools

import (
	"context"
	"mail_service/internal/kit/criteria"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func ParseRequestFilters(ctx context.Context) [][]criteria.Filter {
	request := ctx.Value("request").(*http.Request)
	return parseQueryString(request.URL.RawQuery)
}

func parseQueryString(queryString string) [][]criteria.Filter {
	parsedQuery, err := url.ParseQuery(queryString)
	if err != nil {
		return [][]criteria.Filter{}
	}

	filters := [][]criteria.Filter{}

	for key, values := range parsedQuery {
		for _, value := range values {
			// Separar la clave para obtener índices
			parts := strings.Split(key, "[")
			if len(parts) != 4 {
				return [][]criteria.Filter{}
			}

			// Obtener índices desde las partes
			groupIndex, _ := strconv.Atoi(parts[1])
			criteriaIndex, _ := strconv.Atoi(parts[2])

			// Asegurarse de que el slice tenga suficientes elementos
			for len(filters) <= groupIndex {
				filters = append(filters, []criteria.Filter{})
			}

			for len(filters[groupIndex]) <= criteriaIndex {
				filters[groupIndex] = append(filters[groupIndex], criteria.Filter{})
			}

			// Asignar el valor a la estructura
			switch parts[3] {
			case "campo]":
				filters[groupIndex][criteriaIndex].Campo = value
			case "operador]":
				filters[groupIndex][criteriaIndex].Operador = value
			case "valor]":
				filters[groupIndex][criteriaIndex].Valor = value
			default:
				return [][]criteria.Filter{}
			}

		}
	}

	return filters
}
