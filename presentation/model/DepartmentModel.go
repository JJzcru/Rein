package model

import (
	"Rein/domain"
)

// DepartmentModel Representation of Department in the presentation layer
type DepartmentModel struct {
	DepartmentID string `json:"DepartmentID"`
	Name         string `json:"name"`
	err          domain.Error
}

// NewDepartmentModel Constructor for a new DepartmentModel
func NewDepartmentModel() DepartmentModel {
	return DepartmentModel{}
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

// SetError Set department error
func (d *DepartmentModel) SetError(Error domain.Error) {
	d.err = Error
}

// GetError Get department error
func (d *DepartmentModel) GetError() domain.Error {
	return d.err
}

// ToString Transform a UserModel to string
func (d *DepartmentModel) ToString() string {
	return "\n***** DepartmentModel Details *****\n" +
		"DepartmentID= " + d.GetDepartmentID() + "\n" +
		"name= " + d.GetName() + "\n" +
		"*******************************"
}
