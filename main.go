package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bio426/vugopo/datasource"
	"github.com/bio426/vugopo/route"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	datasource.InitConfig()
	datasource.InitPostgres()

	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	if datasource.Config.PROD {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "./dist",
			HTML5: true,
		}))
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowCredentials: true,
		}))
		// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
		// }))
	}

	route.RegisterApi(e)

	if datasource.Config.PROD {
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
