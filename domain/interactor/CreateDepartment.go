package interactor

import (
	data "Rein/data/repository"
	"Rein/domain"
	"Rein/domain/repository"

	log "github.com/Sirupsen/logrus"
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
	department, err := u.departmentRepository.GetDepartmentByName(u.department.GetName())
	if err != nil {
		department, err = u.departmentRepository.AddDepartment(u.department)
		if err != nil {
			log.Error(err)
			department.SetError(domain.NewError(500, "Error creating the department"))
		}
	} else {
		department.SetError(domain.NewError(409, "Already exist a Department with that name"))
	}

	ch <- department
}
