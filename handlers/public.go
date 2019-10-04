package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/letrannhatviet/my_framework/db"
	"github.com/letrannhatviet/my_framework/db/types"
)

func SearchStudentSimple(c echo.Context) error {
	var req types.StudentReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	student, err := db.GetOneStudent(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func GroupLastName(c echo.Context) error {
	var req types.Student
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	student, err := db.GroupByLast(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func SearchLikeStudent(c echo.Context) error {
	var req types.StudentSearchReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	students, err := db.SearchLikeStudent(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, students)
}

func SearchStudent(c echo.Context) error {
	var req types.StudentReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	students, err := db.SearchStudent(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, students)
}

func GetAllStudent(c echo.Context) error {
	student, _ := db.GetAllStudents()
	return c.JSON(http.StatusOK, student)
}

func GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	res, err := db.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
