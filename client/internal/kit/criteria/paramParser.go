package criteria

import (
	"net/url"
	"strconv"
	"strings"
)

func ParseToFilterArrays(params url.Values) [][]Filter {
	var conditionArrays [][]Filter

	// Iterate over query parameters
	for key, values := range params {
		for _, value := range values {
			// Split key into parts to extract array indices
			parts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(key, "filter["), "]"), "][")

			// Extract array indices and create Filter instances
			index1, _ := strconv.Atoi(parts[0])
			index2, _ := strconv.Atoi(parts[1])

			// Ensure the array is large enough
			for len(conditionArrays) <= index1 {
				conditionArrays = append(conditionArrays, nil)
			}
			for len(conditionArrays[index1]) <= index2 {
				conditionArrays[index1] = append(conditionArrays[index1], Filter{})
			}

			// Set Filter values
			conditionArrays[index1][index2].Campo = getParamValue(value, "campo")
			conditionArrays[index1][index2].Operador = getParamValue(value, "operador")
			conditionArrays[index1][index2].Valor = getParamValue(value, "valor")
		}
	}

	return conditionArrays
}

func getParamValue(value, paramName string) string {
	param := paramName + "="
	startIndex := strings.Index(value, param)
	if startIndex == -1 {
		return ""
	}

	startIndex += len(param)
	endIndex := strings.Index(value[startIndex:], "&")
	if endIndex == -1 {
		endIndex = len(value)
	} else {
		endIndex += startIndex
	}

	return value[startIndex:endIndex]
}
