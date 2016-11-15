package domain

// Department Representation of a Department
type Department struct {
	departmentID string
	name         string
	err          Error
}

// NewDepartment Constructor for a new DepartmentModel
func NewDepartment() Department {
	return Department{}
}

// SetDepartmentID Set the id of the Department
func (d *Department) SetDepartmentID(DepartmentID string) {
	d.departmentID = DepartmentID
}

// GetDepartmentID Get the id of the Department
func (d *Department) GetDepartmentID() string {
	return d.departmentID
}

// SetName Set Department name
func (d *Department) SetName(name string) {
	d.name = name
}

// GetName Get Department name
func (d *Department) GetName() string {
	return d.name
}

// SetError Set an error for the Department
func (d *Department) SetError(err Error) {
	d.err = err
}

// GetError Get an error for the Department
func (d *Department) GetError() Error {
	return d.err
}

// ToString Transform a User to string
func (d *Department) ToString() string {
	return "\n***** Department Details *****\n" +
		"DepartmentID= " + d.GetDepartmentID() + "\n" +
		"name= " + d.GetName() + "\n" +
		"*******************************"
}
