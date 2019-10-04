package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/letrannhatviet/my_framework/db"
	"github.com/letrannhatviet/my_framework/db/types"
)

func AddStudent(c echo.Context) error {
	var req types.StudentAddReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "Bad request", Message: "Bad parameter"})
	}

	i := db.InsertStudent(req)

	return c.JSON(http.StatusOK, i)
}

func DeleteStudent(c echo.Context) error {
	var req types.DeleteReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}

	res, err := db.DeleteStudent(req.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	res, err := db.DeleteStudent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{Code: "BadRequest", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
