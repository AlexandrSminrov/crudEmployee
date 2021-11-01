package service

import (
	"context"
	"crudEmployee/internal/repository"
	mockRepository "crudEmployee/internal/repository/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_AddEmployee(t *testing.T) {

	srv := NewService(mockRepository.NewMockRepository(gomock.NewController(t)))

	_, err := srv.AddEmployee(context.Background(), []*repository.Employee{
		{
			LastName: "test",
		},
	})
	assert.Equal(t, fmt.Errorf("lastName ERROR entry number 1 "), err)
	//assert.Equal(t, test.expectedResponseBody, w.Body.String())
}

func TestService_Update(t *testing.T) {
	srv := NewService(mockRepository.NewMockRepository(gomock.NewController(t)))

	_, err := srv.AddEmployee(context.Background(), []*repository.Employee{
		{
			FirstName: "test",
		},
	})
	assert.Equal(t, fmt.Errorf("firstName ERROR entry number 1 "), err)
}

func TestService_validate(t *testing.T) {
	tests := []struct {
		Name  string
		Out   error
		Input *repository.Employee
	}{
		{
			Name: "First Name err",
			Input: &repository.Employee{
				FirstName: "test",
			},
			Out: fmt.Errorf("firstName ERROR"),
		},
		{
			Name: "Last Name err",
			Input: &repository.Employee{
				LastName: "test",
			},
			Out: fmt.Errorf("lastName ERROR"),
		},
		{
			Name: "Middle Name err",
			Input: &repository.Employee{
				MiddleName: "test",
			},
			Out: fmt.Errorf("middleName ERROR"),
		},
		{
			Name: "Date err",
			Input: &repository.Employee{
				DateOfBirth: "2002-10-32",
			},
			Out: fmt.Errorf("date ERROR "),
		},
		{
			Name: "Address err",
			Input: &repository.Employee{
				Address: "test  :;;;",
			},
			Out: fmt.Errorf("address ERROR "),
		},
		{
			Name: "Department err",
			Input: &repository.Employee{
				Department: "test :;;;",
			},
			Out: fmt.Errorf("department ERROR "),
		},
		{
			Name: "AboutMe err",
			Input: &repository.Employee{
				AboutMe: "test :;;;",
			},
			Out: fmt.Errorf("aboutMe ERROR"),
		},
		{
			Name: "phone err",
			Input: &repository.Employee{
				Phone: "123123P",
			},
			Out: fmt.Errorf("phone number ERROR "),
		},
		{
			Name: "Email err",
			Input: &repository.Employee{
				Email: "test@lklkl",
			},
			Out: fmt.Errorf("email ERROR "),
		},
		{
			Name: "OK",
			Input: &repository.Employee{
				ID:          1,
				FirstName:   "Иван",
				LastName:    "Иванов",
				MiddleName:  "Иванович",
				DateOfBirth: "1980-10-20",
				Address:     "Москва",
				Department:  "тест",
				AboutMe:     "тест",
				Phone:       "77777777",
				Email:       "test@test.ru",
			},
			Out: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := validate(test.Input)
			assert.Equal(t, test.Out, err)
		})
	}

}
