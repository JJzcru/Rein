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
func (d *DepartmentDataRepository) AddDepartment(dep domain.Department) <-chan domain.Department {
	c := make(chan domain.Department, 1)
	//c <- department
	department, err := d.departmentDataStoreFactory.Create().DepartmentCreate(dep)

	if err != nil {
		department.SetError(domain.NewError(500, "Error creating the Department"))
	}
	c <- department
	return c
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
