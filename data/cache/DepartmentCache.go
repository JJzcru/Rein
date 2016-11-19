package cache

import "Rein/domain"

// DepartmentCache Interface for the creation of a cache for Department
type DepartmentCache interface {
	Get(DepartmentID string) (domain.Department, error)
	GetAll() ([]domain.Department, error)
	GetByName(name string) (domain.Department, error)
	Create(department domain.Department) (domain.Department, error)
	Update(department domain.Department) (domain.Department, error)
	Delete(DepartmentID string) (domain.Department, error)
}
