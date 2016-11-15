package repository

import (
	"Rein/domain"
)

// DepartmentRepository handles the repository data for the Department
type DepartmentRepository interface {
	AddDepartment(department domain.Department) <-chan domain.Department
	UpdateDepartment(department domain.Department) domain.Department
	DeleteDepartment(DepartmentID string) domain.Department
	GetDepartment(DepartmentID string) domain.Department
	GetDepartmentByName(name string) domain.Department
	GetDepartments() <-chan []domain.Department
}
