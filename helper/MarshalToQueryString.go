package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
)

func MarshalToQueryString(in any) (result string, err error) {
	bytes, err := json.Marshal(in)
	if err != nil {
		return
	}

	normalized := make(map[string]any)
	err = json.Unmarshal(bytes, &normalized)
	if err != nil {
		return
	}

	queryParts := make([]string, 0, len(normalized))
	for key, value := range normalized {
		part := fmt.Sprintf("%s=", key)

		if value == nil {
			continue
		}

		if _, ok := value.(bool); ok {
			part += fmt.Sprintf("%t", value)
		} else if _, ok := value.(string); ok {
			part += value.(string)
		} else if _, ok := value.(int); ok {
			part += fmt.Sprintf("%d", value)
		} else if _, ok := value.(uint); ok {
			part += fmt.Sprintf("%d", value)
		} else if _, ok := value.(float64); ok {
			if value.(float64) == math.Trunc(value.(float64)) {
				part += fmt.Sprintf("%d", int64(value.(float64)))
			} else {
				part += fmt.Sprintf("%f", value)
			}
		} else {
			err = errors.New(fmt.Sprintf("Unknown type: %T", value))
			return
		}

		queryParts = append(queryParts, part)
	}

	return strings.Join(queryParts, "&"), nil
}
