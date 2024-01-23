package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/vugopo/pkg/auth"
	"github.com/bio426/vugopo/pkg/task"
)

func RegisterApi(app *echo.Echo) {
	apiGroup := app.Group("/api")

	authGoup := apiGroup.Group("/auth")
	authGoup.POST("/register", auth.Controller.Register)
	authGoup.POST("/login", auth.Controller.Login)
	authGoup.POST("/logout", auth.Controller.Logout)

	taskGroup := apiGroup.Group("/task")
	taskGroup.Use(auth.MiddlewareWithSkipper([]string{"/public"}))
	taskGroup.GET("/public", task.Controller.ListPublic)
	taskGroup.GET("", task.Controller.ListByUser)
	taskGroup.POST("", task.Controller.Create)

	// SSE
	sseGroup := app.Group("/event")
	sseGroup.GET("/test", task.Controller.MyTest)
}
