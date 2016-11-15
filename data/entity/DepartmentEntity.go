package entity

// DepartmentEntity Representation of a Department in the data layer
type DepartmentEntity struct {
	DepartmentID string `json:"DepartmentID" bson:"DepartmentID"`
	Name         string `json:"name" bson:"name"`
}

// NewDepartmentEntity Constructor for a new DepartmentModel
func NewDepartmentEntity() DepartmentEntity {
	return DepartmentEntity{}
}

// SetDepartmentID Set the id of the department
func (d *DepartmentEntity) SetDepartmentID(DepartmentID string) {
	d.DepartmentID = DepartmentID
}

// GetDepartmentID Get the id of the department
func (d *DepartmentEntity) GetDepartmentID() string {
	return d.DepartmentID
}

// SetName Set department name
func (d *DepartmentEntity) SetName(name string) {
	d.Name = name
}

// GetName Get department name
func (d *DepartmentEntity) GetName() string {
	return d.Name
}

// ToString Transform a User to string
func (d *DepartmentEntity) ToString() string {
	return "\n***** DepartmentEntity Details *****\n" +
		"DepartmentID= " + d.GetDepartmentID() + "\n" +
		"name= " + d.GetName() + "\n" +
		"*******************************"
}
