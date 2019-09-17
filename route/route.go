package route

import (
	"github.com/labstack/echo"
	"github.com/letrannhatviet/my_framework/handlers"
)

func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")
	publicRoute.GET("/student", handlers.GetAllStudent)
	publicRoute.PATCH("/student", handlers.GetStudent)
	publicRoute.GET("/health", handlers.CheckHealth)
}

func Staff(e *echo.Echo) {
	staffRoute := e.Group("/v1/staff")
	staffRoute.POST("/student", handlers.AddStudent)
	staffRoute.DELETE("/student", handlers.DeleteStudent)
}
