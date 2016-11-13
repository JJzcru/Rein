package model

// DepartmentModel Representation of Department in the presentation layer
type DepartmentModel struct {
	DepartmentID string `json:"DepartmentID,omitempty"`
	Name         string `json:"name,omitempty"`
}

// NewDepartmentModel Constructor for a new DepartmentModel
func NewDepartmentModel() *DepartmentModel {
	return &DepartmentModel{}
}

// SetDepartmentID Set the id of the department
func (d *DepartmentModel) SetDepartmentID(DepartmentID string) {
	d.DepartmentID = DepartmentID
}

// GetDepartmentID Get the id of the department
func (d *DepartmentModel) GetDepartmentID() string {
	return d.DepartmentID
}

// SetName Set department name
func (d *DepartmentModel) SetName(name string) {
	d.Name = name
}

// GetName Get department name
func (d *DepartmentModel) GetName() string {
	return d.Name
}

// ToString Transform a UserModel to string
func (d *DepartmentModel) ToString() string {
	return "\n***** DepartmentModel Details *****\n" +
		"DepartmentID= " + d.GetDepartmentID() + "\n" +
		"name= " + d.GetName() + "\n" +
		"*******************************"
}
