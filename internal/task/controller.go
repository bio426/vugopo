package task

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/vugopo/internal/core"
)

type AuthCtl core.Controller

type CtlListRow struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Category *string `json:"category,omitempty"`
}
type CtlListResponse struct {
	Rows []CtlListRow `json:"rows"`
}

func (ctl *AuthCtl) List(c echo.Context) error {
	res, err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

var Controller = &AuthCtl{}
