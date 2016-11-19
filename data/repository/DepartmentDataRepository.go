package repository

import (
	"Rein/data/repository/datasource"
	"Rein/domain"
	"Rein/domain/repository"
)

// DepartmentDataRepository Implementation of the DepartmentRepository
type DepartmentDataRepository struct {
	repository.DepartmentRepository
	departmentDataStoreFactory datasource.DatabaseDepartmentDataStoreFactory
}

// NewDepartmentDataRepository Constructor for a new DepartmentDataRepository
func NewDepartmentDataRepository() *DepartmentDataRepository {
	return &DepartmentDataRepository{departmentDataStoreFactory: datasource.NewDatabaseDepartmentDataStoreFactory()}
}

// AddDepartment Add a department to the repository
func (d *DepartmentDataRepository) AddDepartment(dep domain.Department) (domain.Department, error) {
	return d.departmentDataStoreFactory.Create().DepartmentCreate(dep)
}

// GetDepartments Get a list of all the departments
func (d *DepartmentDataRepository) GetDepartments() <-chan []domain.Department {
	c := make(chan []domain.Department, 1)
	//c <- department
	departments, err := d.departmentDataStoreFactory.Create().DepartmentGetAll()
	if err != nil {
		for i := range departments {
			departments[i].SetError(domain.NewError(500, "Error getting the Departments"))
		}
	}
	c <- departments
	return c
}

// GetDepartmentByName return a list of all the departments by name
func (d *DepartmentDataRepository) GetDepartmentByName(name string) (domain.Department, error) {
	return d.departmentDataStoreFactory.Create().DepartmentGetByName(name)
}
