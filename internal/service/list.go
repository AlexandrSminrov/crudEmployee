package service

import (
	"context"
	"crudEmployee/internal/repository"
	"fmt"
	"time"

	joiner "github.com/json-iterator/go"
)

// GetAll get all employees from db employee
func (s *service) GetAll(ctx context.Context) ([]byte, error) {
	r, err := s.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return joiner.Marshal(r)
}

// AddEmployee add records to the database
func (s *service) AddEmployee(ctx context.Context, employees []*repository.Employee) (string, error) {
	for i, employee := range employees {
		if err := validate(employee); err != nil {
			return "", fmt.Errorf("%v entry number %d ", err, i)
		}
	}

	rowsAffected, err := s.Repository.AddEmployee(ctx, employees)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("entries added %d", rowsAffected), err
}

func (s *service) GetByID(ctx context.Context, id int64) ([]byte, error) {
	employee, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return joiner.Marshal(employee)
}

func (s *service) Update(ctx context.Context, id int64, employee *repository.Employee) (string, error) {
	if err := validate(employee); err != nil {
		return "", fmt.Errorf("%v entry number", err)
	}

	rowsAffected, err := s.Repository.Update(ctx, id, employee)
	return fmt.Sprintf("update rows: %d", rowsAffected), err
}

// validate verifies the request
func validate(employee *repository.Employee) error {
	if onlyRu.MatchString(employee.FirstName) {
		return fmt.Errorf("firstName ERROR")
	}
	if onlyRu.MatchString(employee.LastName) {
		return fmt.Errorf("lastName ERROR")
	}
	if onlyRu.MatchString(employee.MiddleName) {
		return fmt.Errorf("middleName ERROR")
	}
	if _, err := time.Parse("2006-01-02", employee.DateOfBirth); err != nil && len(employee.DateOfBirth) > 1 {
		return fmt.Errorf("date ERROR ")
	}
	if address.MatchString(employee.Address) {
		return fmt.Errorf("address ERROR ")
	}
	if onlyRuEng.MatchString(employee.Department) {
		return fmt.Errorf("department ERROR ")
	}
	if address.MatchString(employee.AboutMe) {
		return fmt.Errorf("aboutMe ERROR")
	}
	if onlyNum.MatchString(employee.Phone) {
		return fmt.Errorf("phone number ERROR ")
	}
	if !email.MatchString(employee.Email) && len(employee.Email) > 1 {
		return fmt.Errorf("email ERROR ")
	}

	return nil
}
