package route

import (
	"github.com/labstack/echo"
	"github.com/letrannhatviet/my_framework/handlers"
)

func Public(e *echo.Echo) {
	publicRoute := e.Group("/api/v1/public")
	publicRoute.GET("/student", handlers.GetAllStudent)
	publicRoute.GET("/student/id/:id", handlers.GetById)
	publicRoute.GET("/student/group/last_name", handlers.GroupLastName)
	publicRoute.PATCH("/student/simple", handlers.SearchStudentSimple)
	publicRoute.PATCH("/student", handlers.SearchStudent)
	publicRoute.PATCH("/student1", handlers.SearchLikeStudent)
	publicRoute.GET("/health", handlers.CheckHealth)
}

func Staff(e *echo.Echo) {
	staffRoute := e.Group("/api/v1/staff")
	staffRoute.POST("/student", handlers.AddStudent)
	staffRoute.DELETE("/student", handlers.DeleteStudent)
	staffRoute.DELETE("/student/id/:id", handlers.DeleteById)
}
