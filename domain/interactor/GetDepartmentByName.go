package interactor

import (
	data "Rein/data/repository"
	"Rein/domain"
	"Rein/domain/repository"
)

// GetDepartmentByName Use case to get a Department by its name
type GetDepartmentByName struct {
	UseCase
	name                 string
	departmentRepository repository.DepartmentRepository
}

// NewGetDepartmentByName Constructor for CreateDepartment UseCase
func NewGetDepartmentByName(name string) *GetDepartmentByName {
	return &GetDepartmentByName{departmentRepository: data.NewDepartmentDataRepository(), name: name}
}

// Execute The use case
func (u *GetDepartmentByName) Execute(ch chan domain.Department) {
	department, err := u.departmentRepository.GetDepartmentByName(u.name)
	if err != nil {
		department.SetError(domain.NewError(500, "Error creating the department"))
	}

	if department.GetName() == "" {
		department.SetError(domain.NewError(404, "The department with name '"+u.name+"' don't exist"))
	}
	ch <- department
}
