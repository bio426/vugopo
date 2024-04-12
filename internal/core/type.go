package core

import (
	"database/sql"

	"github.com/bio426/vugopo/datasource"
)

type Controller struct {
	Svc *Service
}

type Service struct {
	Config *datasource.Configuration
	Db     *sql.DB
}

type PaginatedResponse[T any] struct {
	Total int32 `json:"total"`
	From  int32 `json:"from"`
	To    int32 `json:"to"`
	Rows  []T   `json:"rows"`
}
