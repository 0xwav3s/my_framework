package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letrannhatviet/my_framework/db"
	"github.com/letrannhatviet/my_framework/db/types"
)

func GetStudent(c echo.Context) error {
	var req types.StudentReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: "Bad request"})
	}
	student, err := db.GetStudent()
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func GetAllStudent(c echo.Context) error {
	student, _ := db.GetAllStudent()
	return c.JSON(http.StatusOK, student)
}

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
