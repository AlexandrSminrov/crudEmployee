package repository

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// GetAll will return all employees from the table
func (r *repository) GetAll(ctx context.Context) ([]Employee, error) {
	query := `SELECT id, first_name, last_name, middle_name, date_of_birth, address, department, about_me, phone, email  FROM public.employees`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("db query error")
	}
	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("error rows.Close %v", err.Error())
		}
	}()
	var employees []Employee

	for rows.Next() {
		var employee Employee
		if err = rows.Scan(
			&employee.ID,
			&employee.FirstName,
			&employee.LastName,
			&employee.MiddleName,
			&employee.DateOfBirth,
			&employee.Address,
			&employee.Department,
			&employee.AboutMe,
			&employee.Phone,
			&employee.Email,
		); err != nil {
			return nil, fmt.Errorf("db scan rows error")
		}
		employees = append(employees, employee)
	}
	fmt.Println(employees)
	return employees, nil
}

// AddEmployee add employees to the table
func (r *repository) AddEmployee(ctx context.Context, employees []*Employee) (int64, error) {
	insert := sq.Insert("public.employees").Columns(
		"first_name",
		"last_name",
		"middle_name",
		"date_of_birth",
		"address",
		"department",
		"about_me",
		"phone",
		"email",
	)

	for _, employee := range employees {
		insert = insert.Values(
			employee.FirstName,
			employee.LastName,
			employee.MiddleName,
			employee.DateOfBirth,
			employee.Address,
			employee.Department,
			employee.AboutMe,
			employee.Phone,
			employee.Email,
		)
	}

	query, args, err := insert.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error in the formation of the sql query")
	}

	resExec, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("sql query execution error")
	}
	return resExec.RowsAffected()
}

// GetByID employee request by id
func (r *repository) GetByID(ctx context.Context, id int64) (*Employee, error) {
	query := `SELECT id, first_name, last_name, middle_name, date_of_birth, address, department, about_me, phone, email FROM public.employees WHERE id in ($1)`

	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("error sql query %v", err)
	}
	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("error rows.Close %v", err.Error())
		}
	}()

	var employee Employee

	for rows.Next() {
		if err = rows.Scan(
			&employee.ID,
			&employee.FirstName,
			&employee.LastName,
			&employee.MiddleName,
			&employee.DateOfBirth,
			&employee.Address,
			&employee.Department,
			&employee.AboutMe,
			&employee.Phone,
			&employee.Email,
		); err != nil {
			return nil, err
		}
	}
	if employee.ID == 0 {
		return nil, fmt.Errorf("employee not fond")
	}
	return &employee, nil
}

func (r *repository) Update(ctx context.Context, id int64, employee *Employee) (int64, error) {
	sqlReq := sq.Update("public.employees")

	v := reflect.ValueOf(*employee)
	typeOfS := v.Type()

	for i := 1; i < v.NumField(); i++ {
		val := fmt.Sprint(v.Field(i).Interface())
		if len(val) > 0 {
			if nameColumn, ok := mapNameColumn[typeOfS.Field(i).Name]; ok {
				sqlReq = sqlReq.Set(nameColumn, val)
			}
		}
	}

	query, args, err := sqlReq.Where("id in (?)", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error in the formation of the sql query")
	}

	rows, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("sql query execution error")
	}

	rowsAffected, err := rows.RowsAffected()
	if rowsAffected == 0 {
		return rowsAffected, fmt.Errorf("employee not fond")
	}

	return rowsAffected, err
}
