package handler

import (
	"crudEmployee/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler methods
type Handler interface {
	GetEmployees(c echo.Context) error
	AddEmployees(c echo.Context) error
	GetByIDEmployee(c echo.Context) error
	UpdateEmployee(c echo.Context) error
}

// Router struct
type handler struct {
	services service.Service
}

// NewHandler returns handler
func NewHandler(services service.Service) Handler {
	return &handler{services: services}
}

// InitHandler initializes the routers
func InitHandler(h Handler) *echo.Echo {
	s := echo.New()

	routers := []struct {
		Name        string
		Method      string
		Path        string
		HandlerFunc echo.HandlerFunc
	}{
		{
			"GetEmployees",
			http.MethodGet,
			"/employee",
			h.GetEmployees,
		},
		{
			"AddEmployee",
			http.MethodPost,
			"/employee",
			h.AddEmployees,
		},
		{
			"GetEmployee",
			http.MethodGet,
			"/employee/:id",
			h.GetByIDEmployee,
		},
		{
			"UpEmployee",
			http.MethodPut,
			"/employee/:id",
			h.UpdateEmployee,
		},
	}

	for _, router := range routers {
		s.Router().Add(router.Method, router.Path, router.HandlerFunc)
	}

	s.Use(middleware.Logger())
	return s
}
