package domain

// Department Representation of a Department
type Department struct {
	departmentID string
	name         string
}

// NewDepartment Constructor for a new DepartmentModel
func NewDepartment() *Department {
	return &Department{}
}

// SetDepartmentID Set the id of the department
func (d *Department) SetDepartmentID(DepartmentID string) {
	d.departmentID = DepartmentID
}

// GetDepartmentID Get the id of the department
func (d *Department) GetDepartmentID() string {
	return d.departmentID
}

// SetName Set department name
func (d *Department) SetName(name string) {
	d.name = name
}

// GetName Get department name
func (d *Department) GetName() string {
	return d.name
}

// ToString Transform a User to string
func (d *Department) ToString() string {
	return "\n***** Department Details *****\n" +
		"DepartmentID= " + d.GetDepartmentID() + "\n" +
		"name= " + d.GetName() + "\n" +
		"*******************************"
}
