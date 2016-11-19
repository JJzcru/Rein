package cache

import (
	"Rein/domain"

	uuid "github.com/nu7hatch/gouuid"

	mapper "Rein/data/entity/mapper"

	"Rein/data/entity"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DepartmentCacheImpl Implementation of the DepartmentRepository
type DepartmentCacheImpl struct {
	DepartmentCache
	departmentEntityDataMapper *mapper.DepartmentEntityDataMapper
}

// NewDepartmentCacheImpl Constructor for a new Department Cache implementation
func NewDepartmentCacheImpl() *DepartmentCacheImpl {
	return &DepartmentCacheImpl{departmentEntityDataMapper: mapper.NewDepartmentEntityDataMapper()}
}

// Create return the created Departmen
func (d *DepartmentCacheImpl) Create(department domain.Department) (domain.Department, error) {
	session, err := mgo.Dial("db")
	if err != nil {
		return department, err
	}

	defer session.Close()
	c := session.DB("r-manager").C("department")

	id, err := uuid.NewV4()
	if err != nil {
		return department, err
	}

	department.SetDepartmentID(id.String())

	departmentEntity := d.departmentEntityDataMapper.Convert(department)
	err = c.Insert(departmentEntity)
	if err != nil {
		panic(err)
	}
	return department, err
}

// GetAll return a list of all the departments
func (d *DepartmentCacheImpl) GetAll() ([]domain.Department, error) {
	var departments []domain.Department
	var departmentsEntities []entity.DepartmentEntity

	session, err := mgo.Dial("db")
	if err != nil {
		return departments, err
	}

	defer session.Close()
	c := session.DB("r-manager").C("department")

	err = c.Find(nil).All(&departmentsEntities)
	departments = d.departmentEntityDataMapper.TransformSlice(departmentsEntities)
	return departments, err
}

// GetByName return a department by the name
func (d *DepartmentCacheImpl) GetByName(name string) (domain.Department, error) {
	var department domain.Department
	var departmentEntity entity.DepartmentEntity

	session, err := mgo.Dial("db")
	if err != nil {
		return department, err
	}

	defer session.Close()
	c := session.DB("r-manager").C("department")

	err = c.Find(bson.M{"name": name}).One(&departmentEntity)
	department = d.departmentEntityDataMapper.Transform(departmentEntity)
	return department, err
}
