package auth

import (
	// "github.com/bio426/vugopo/datasource"
	"github.com/bio426/vugopo/pkg/core"
)

type AuthSvc core.Service

func (svc *AuthSvc) GenerateToken() error {
	// datasource.Postgres.ExecContext()

	return nil
}

var Service = &AuthSvc{}
