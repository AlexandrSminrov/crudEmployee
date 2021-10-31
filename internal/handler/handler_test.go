package handler

import (
	"bytes"
	"context"
	"crudEmployee/internal/repository"
	mockRepository "crudEmployee/internal/repository/mocks"
	"crudEmployee/internal/service"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_AddEmployees(t *testing.T) {
	type mockBeaver func(m *mockRepository.MockRepository)

	tests := []struct {
		Name                 string
		InputHeaders         map[string]string
		mock                 bool
		mockBeaver           mockBeaver
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			Name: "OK",
			InputHeaders: map[string]string{
				"Accept": "application/json",
			},
			mock: true,
			mockBeaver: func(m *mockRepository.MockRepository) {
				m.EXPECT().GetAll(context.Background()).Return([]repository.Employee{
					{
						ID:          1,
						FirstName:   "Иван",
						LastName:    "Иванов",
						MiddleName:  "Иванович",
						DateOfBirth: "1980-10-11T00:00:00Z",
						Address:     "Москва",
						Department:  "тест",
						AboutMe:     "тест",
						Phone:       "77777777",
						Email:       "test@test.ru",
					},
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"first_name":"Иван","last_name":"Иванов","middle_name":"Иванович","date_of_birth":"1980-10-11T00:00:00Z","address":"Москва","department":"тест","about_me":"тест","phone":"77777777","email":"test@test.ru"}]`,
		},
		{
			Name:                 "error header",
			mock:                 false,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `header accept error`,
		},
		{
			Name: "db error",
			InputHeaders: map[string]string{
				"Accept": "application/json",
			},
			mock: true,
			mockBeaver: func(m *mockRepository.MockRepository) {
				m.EXPECT().GetAll(context.Background()).Return(nil, fmt.Errorf("db query error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `db query error`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepository.NewMockRepository(c)
			handle := NewHandler(service.NewService(repo))
			if test.mock {
				test.mockBeaver(repo)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/employee", nil)

			for val, key := range test.InputHeaders {
				req.Header.Add(val, key)
			}

			_ = handle.GetEmployees(echo.New().NewContext(req, w))

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_AddEmployees2(t *testing.T) {
	type mockBeaver func(m *mockRepository.MockRepository, employees []*repository.Employee)

	tests := []struct {
		Name                 string
		InputBody            string
		mock                 bool
		mockBeaver           mockBeaver
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			Name:                 "OK",
			InputBody:            `[{"first_name":"Иван","last_name":"Иванов","middle_name":"Иванович","date_of_birth":"1980-10-11","address":"Москва","department":"тест","about_me":"тест","phone":"77777777","email":"test@test.ru"}]`,
			mock:                 true,
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `entries added 1`,
			mockBeaver: func(m *mockRepository.MockRepository, employees []*repository.Employee) {
				m.EXPECT().AddEmployee(context.Background(), employees).Return(int64(1), nil)
			},
		},
		{
			Name:                 "Bed JSON",
			InputBody:            `[{first_name":"Иван","last_name":"Иванов","middle_name":"Иванович","date_of_birth":"1980-10-11","address":"Москва","department":"тест","about_me":"тест","phone":"77777777","email":"test@test.ru"}]`,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `error json`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepository.NewMockRepository(c)
			handle := NewHandler(service.NewService(repo))
			if test.mock {
				var employees []*repository.Employee
				assert.Equal(t, nil, jsoniter.Unmarshal([]byte(test.InputBody), &employees))
				test.mockBeaver(repo, employees)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewBufferString(test.InputBody))

			_ = handle.AddEmployees(echo.New().NewContext(req, w))

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetByIDEmployee(t *testing.T) {
	type mockBeaver func(m *mockRepository.MockRepository, id int64)

	tests := []struct {
		Name                 string
		ID                   string
		mock                 bool
		mockBeaver           mockBeaver
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			Name: "OK",
			ID:   "1",
			mock: true,
			mockBeaver: func(m *mockRepository.MockRepository, id int64) {
				m.EXPECT().GetByID(context.Background(), id).Return(&repository.Employee{
					ID:          1,
					FirstName:   "Иван",
					LastName:    "Иванов",
					MiddleName:  "Иванович",
					DateOfBirth: "1980-10-11T00:00:00Z",
					Address:     "Москва",
					Department:  "тест",
					AboutMe:     "тест",
					Phone:       "77777777",
					Email:       "test@test.ru",
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":1,"first_name":"Иван","last_name":"Иванов","middle_name":"Иванович","date_of_birth":"1980-10-11T00:00:00Z","address":"Москва","department":"тест","about_me":"тест","phone":"77777777","email":"test@test.ru"}`,
		},
		{
			Name:                 "Invalid id",
			ID:                   "test",
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `invalid id`,
		},
		{
			Name: "Not found",
			ID:   "2",
			mock: true,
			mockBeaver: func(m *mockRepository.MockRepository, id int64) {
				m.EXPECT().GetByID(context.Background(), id).Return(nil, fmt.Errorf("employee not fond"))
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: "employee not fond",
		},
		{
			Name: "sql error",
			ID:   "2",
			mock: true,
			mockBeaver: func(m *mockRepository.MockRepository, id int64) {
				m.EXPECT().GetByID(context.Background(), id).Return(nil, fmt.Errorf("sql error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: "sql error",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepository.NewMockRepository(c)
			handle := NewHandler(service.NewService(repo))
			if test.mock {
				id, err := strconv.ParseInt(test.ID, 10, 64)
				assert.Equal(t, nil, err)
				test.mockBeaver(repo, id)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/employee/"+test.ID, nil)

			ctx := echo.New().NewContext(req, w)
			ctx.SetParamNames("id")
			ctx.SetParamValues(test.ID)
			_ = handle.GetByIDEmployee(ctx)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})

	}

}

func TestHandler_UpdateEmployee(t *testing.T) {
	type mockBeaver func(m *mockRepository.MockRepository, id int64, employee *repository.Employee)

	tests := []struct {
		Name                 string
		ID                   string
		InputBody            string
		mock                 bool
		mockBeaver           mockBeaver
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			Name:      "OK",
			ID:        "1",
			InputBody: `{"first_name":"Иван"}`,
			mock:      true,
			mockBeaver: func(m *mockRepository.MockRepository, id int64, employee *repository.Employee) {
				m.EXPECT().Update(context.Background(), id, employee).Return(int64(1), nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `update rows: 1`,
		},
		{
			Name:                 "invalid id",
			ID:                   "test",
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `invalid id`,
		},
		{
			Name:                 "invalid JSON",
			ID:                   "1",
			InputBody:            `{"first_name":"Иван}`,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `error json`,
		},
		{
			Name:      "not found",
			ID:        "1",
			InputBody: `{"first_name":"Иван"}`,
			mock:      true,
			mockBeaver: func(m *mockRepository.MockRepository, id int64, employee *repository.Employee) {
				m.EXPECT().Update(context.Background(), id, employee).Return(int64(0), fmt.Errorf("employee not fond"))
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `employee not fond`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockRepository.NewMockRepository(c)
			handle := NewHandler(service.NewService(repo))
			if test.mock {
				id, err := strconv.ParseInt(test.ID, 10, 64)
				assert.Equal(t, nil, err)

				var employee *repository.Employee
				assert.Equal(t, nil, jsoniter.Unmarshal([]byte(test.InputBody), &employee))
				test.mockBeaver(repo, id, employee)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/employee/"+test.ID, bytes.NewBufferString(test.InputBody))

			ctx := echo.New().NewContext(req, w)
			ctx.SetParamNames("id")
			ctx.SetParamValues(test.ID)
			_ = handle.UpdateEmployee(ctx)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})

	}

}
