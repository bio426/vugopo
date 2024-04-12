package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/vugopo/internal/auth"
	"github.com/bio426/vugopo/internal/task"
)

func RegisterApi(app *echo.Echo) {
	apiGroup := app.Group("/api")

	authGoup := apiGroup.Group("/auth")
	// authGoup.POST("/register", auth.Controller.Register)
	authGoup.POST("/login", auth.Controller.Login)
	authGoup.POST("/logout", auth.Controller.Logout)

	taskGroup := apiGroup.Group("/task")
	taskGroup.GET("", task.Controller.List)
	taskGroup.POST("", task.Controller.List)

	// SSE
	// sseGroup := app.Group("/event")
	// sseGroup.GET("/test", task.Controller.MyTest)
}
