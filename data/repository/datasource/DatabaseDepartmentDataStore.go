package datasource

import (
	"Rein/data/cache"
	"Rein/domain"
)

// DatabaseDepartmentDataStore Implementation based on database system data store.
type DatabaseDepartmentDataStore struct {
	DepartmentDataStore
	departmentCache cache.DepartmentCache
}

// NewDatabaseDepartmentDataStore Constructor for a new Department Data Store using a database
func NewDatabaseDepartmentDataStore(departmentCache cache.DepartmentCache) *DatabaseDepartmentDataStore {
	return &DatabaseDepartmentDataStore{departmentCache: departmentCache}
}

// DepartmentCreate Create a new Department
func (d *DatabaseDepartmentDataStore) DepartmentCreate(department domain.Department) (domain.Department, error) {
	return d.departmentCache.Create(department)
}

// DepartmentCreate Create a new Department
func (d *DatabaseDepartmentDataStore) DepartmentGetAll() ([]domain.Department, error) {
	return d.departmentCache.GetAll()
}
