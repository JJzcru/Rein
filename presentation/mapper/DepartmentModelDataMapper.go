package mapper

import (
	"Rein/domain"
	"Rein/presentation/model"
)

// DepartmentModelDataMapper Mapper that transform a departmentModel to a department
type DepartmentModelDataMapper struct{}

// NewDepartmentModelDataMapperr Constructor for the mapper
func NewDepartmentModelDataMapperr() *DepartmentModelDataMapper {
	d := new(DepartmentModelDataMapper)
	return d
}

// Transform a DepartmentModel to Department
func (d *DepartmentModelDataMapper) Transform(departmentModel model.DepartmentModel) *domain.Department {
	deparment := domain.NewDepartment()
	deparment.SetDepartmentID(departmentModel.GetDepartmentID())
	deparment.SetName(departmentModel.GetName())
	return deparment
}

// TransformSlice transform a Slice of departmentEntity to Slice of department
func (d *DepartmentModelDataMapper) TransformSlice(departmentModelSlice []model.DepartmentModel) []*domain.Department {
	departmentSlice := make([]*domain.Department, len(departmentModelSlice))
	for i, departmentModel := range departmentModelSlice {
		departmentSlice[i] = d.Transform(departmentModel)
	}
	return departmentSlice
}

// Convert a department to departmentEntity
func (d *DepartmentModelDataMapper) Convert(department domain.Department) *model.DepartmentModel {
	departmentModel := model.NewDepartmentModel()
	departmentModel.SetDepartmentID(department.GetDepartmentID())
	departmentModel.SetName(department.GetName())
	return departmentModel
}

// ConvertSlice transform a Slice of department to Slice of departmentEntity
func (d *DepartmentModelDataMapper) ConvertSlice(departmentSlice []domain.Department) []*model.DepartmentModel {
	departmentModelSlice := make([]*model.DepartmentModel, len(departmentSlice))
	for i, department := range departmentSlice {
		departmentModelSlice[i] = d.Convert(department)
	}
	return departmentModelSlice
}
