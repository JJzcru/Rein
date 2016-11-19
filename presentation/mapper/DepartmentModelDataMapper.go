package mapper

import (
	"Rein/domain"
	"Rein/presentation/model"
)

// DepartmentModelDataMapper Mapper that transform a departmentModel to a department
type DepartmentModelDataMapper struct{}

// NewDepartmentModelDataMapper Constructor for the mapper
func NewDepartmentModelDataMapper() *DepartmentModelDataMapper {
	d := new(DepartmentModelDataMapper)
	return d
}

// Transform a DepartmentModel to Department
func (d *DepartmentModelDataMapper) Transform(departmentModel model.DepartmentModel) domain.Department {
	department := domain.NewDepartment()
	department.SetDepartmentID(departmentModel.GetDepartmentID())
	department.SetName(departmentModel.GetName())
	department.SetError(departmentModel.GetError())
	return department
}

// TransformSlice transform a Slice of departmentEntity to Slice of department
func (d *DepartmentModelDataMapper) TransformSlice(departmentModelSlice []model.DepartmentModel) []domain.Department {
	departmentSlice := make([]domain.Department, len(departmentModelSlice))
	for i, departmentModel := range departmentModelSlice {
		departmentSlice[i] = d.Transform(departmentModel)
	}
	return departmentSlice
}

// Convert a department to departmentEntity
func (d *DepartmentModelDataMapper) Convert(department domain.Department) model.DepartmentModel {
	departmentModel := model.NewDepartmentModel()
	departmentModel.SetDepartmentID(department.GetDepartmentID())
	departmentModel.SetName(department.GetName())
	departmentModel.SetError(department.GetError())
	return departmentModel
}

// ConvertSlice transform a Slice of department to Slice of departmentEntity
func (d *DepartmentModelDataMapper) ConvertSlice(departmentSlice []domain.Department) []model.DepartmentModel {
	departmentModelSlice := make([]model.DepartmentModel, len(departmentSlice))
	for i, department := range departmentSlice {
		departmentModelSlice[i] = d.Convert(department)
	}
	return departmentModelSlice
}
