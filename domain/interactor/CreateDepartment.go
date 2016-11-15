package interactor

import (
	data "Rein/data/repository"
	"Rein/domain"
	"Rein/domain/repository"
)

// CreateDepartment Use case for create a new Department
type CreateDepartment struct {
	UseCase
	department           domain.Department
	departmentRepository repository.DepartmentRepository
}

// NewCreateDepartment Constructor for CreateDepartment UseCase
func NewCreateDepartment(department domain.Department) *CreateDepartment {
	return &CreateDepartment{departmentRepository: data.NewDepartmentDataRepository(), department: department}
}

// Execute The use case
func (u *CreateDepartment) Execute(ch chan domain.Department) {
	/*c := u.departmentRepository.AddDepartment(u.department)
	department := <-c
	fmt.Println(department.ToString())

	ch <- department*/

	ch <- <-u.departmentRepository.AddDepartment(u.department)
}
