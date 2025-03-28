package utils

import (
	"errors"
	"net/http"
	"strconv"
)

func GetQueryInt(r *http.Request, key string, defaultValue int) (int, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil || intValue < 0 {
		return 0, errors.New("invalid integer value")
	}
	return intValue, nil
}

func GetQueryString(r *http.Request, key, defaultValue string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetQueryFilters(r *http.Request) map[string]string {
	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		if key != "limit" && key != "offset" && key != "sortBy" && key != "sortOrder" {
			filters[key] = values[0]
		}
	}
	return filters
}
