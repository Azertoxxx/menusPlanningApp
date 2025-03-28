package model

type QueryParams struct {
	Limit     int               `json:"limit"`
	Offset    int               `json:"offset"`
	SortBy    string            `json:"sortBy"`
	SortOrder string            `json:"sortOrder"`
	Filters   map[string]string `json:"filters"`
}
