package datasource

import (
	"Rein/domain"
)

// DepartmentDataStore Handles the department data
type DepartmentDataStore interface {
	DepartmentGetAll() ([]domain.Department, error)
	DepartmentCreate(department domain.Department) (domain.Department, error)
	DepartmentGetByName(name string) (domain.Department, error)
}
