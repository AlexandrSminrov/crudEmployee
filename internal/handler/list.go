package handler

import (
	"crudEmployee/internal/repository"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	joiner "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetEmployees(c echo.Context) error {
	t := c.Request().Header.Get("Accept")
	if strings.Contains(t, "application/json") || strings.Contains(t, "*/*") {
		res, err := h.services.GetAll(c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, res)

	}
	return c.String(http.StatusBadRequest, fmt.Sprint("header accept error"))
}

func (h *handler) AddEmployees(c echo.Context) error {
	var employees []*repository.Employee

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("error body"))
	}

	if err = joiner.Unmarshal(body, &employees); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("error json"))
	}

	response, err := h.services.AddEmployee(c.Request().Context(), employees)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, response)
}

func (h *handler) GetByIDEmployee(c echo.Context) error {
	s := c.Param("id")
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	employee, err := h.services.GetByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "employee not fond" {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, employee)
}

func (h *handler) UpdateEmployee(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	var employees *repository.Employee

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("error body"))
	}

	if err = joiner.Unmarshal(body, &employees); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprint("error json"))
	}

	response, err := h.services.Update(c.Request().Context(), id, employees)
	if err != nil {
		if err.Error() == "employee not fond" {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}
