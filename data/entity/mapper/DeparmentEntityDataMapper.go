package mapper

import (
	"Rein/data/entity"
	"Rein/domain"
)

// DepartmentEntityDataMapper Mapper that transform a DepartmentEntity to a Department
type DepartmentEntityDataMapper struct{}

// NewDepartmentEntityDataMapperr Constructor for the mapper
func NewDepartmentEntityDataMapperr() *DepartmentEntityDataMapper {
	d := new(DepartmentEntityDataMapper)
	return d
}

// Transform a DepartmentEntity to Department
func (d *DepartmentEntityDataMapper) Transform(departmentEntity entity.DepartmentEntity) *domain.Department {
	deparment := domain.NewDepartment()
	deparment.SetDepartmentID(departmentEntity.GetDepartmentID())
	deparment.SetName(departmentEntity.GetName())
	return deparment
}

// TransformSlice transform a Slice of DepartmentEntity to Slice of Department
func (d *DepartmentEntityDataMapper) TransformSlice(departmentEntitySlice []entity.DepartmentEntity) []*domain.Department {
	departmentSlice := make([]*domain.Department, len(departmentEntitySlice))
	for i, departmentEntity := range departmentEntitySlice {
		departmentSlice[i] = d.Transform(departmentEntity)
	}
	return departmentSlice
}

// Convert a Department to DepartmentEntity
func (d *DepartmentEntityDataMapper) Convert(department domain.Department) *entity.DepartmentEntity {
	departmentEntity := entity.NewDepartmentEntity()
	departmentEntity.SetDepartmentID(department.GetDepartmentID())
	departmentEntity.SetName(department.GetName())
	return departmentEntity
}

// ConvertSlice transform a Slice of Department to Slice of DepartmentEntity
func (d *DepartmentEntityDataMapper) ConvertSlice(departmentSlice []domain.Department) []*entity.DepartmentEntity {
	departmentEntitySlice := make([]*entity.DepartmentEntity, len(departmentSlice))
	for i, department := range departmentSlice {
		departmentEntitySlice[i] = d.Convert(department)
	}
	return departmentEntitySlice
}
