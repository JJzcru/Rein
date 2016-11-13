package model

import (
	"strconv"
	"strings"
	"time"
)

// UserModel Representation of User in the presentation layer
type UserModel struct {
	UserID       string    `json:"UserID"`
	DepartmentID string    `json:"DepartmentID"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Role         []string  `json:"role"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// NewUserModel Constructor for a new UserModel
func NewUserModel() *UserModel {
	return &UserModel{}
}

// SetUserID set the id of the user
func (u *UserModel) SetUserID(UserID string) {
	u.UserID = UserID
}

// GetUserID get the id of the user
func (u *UserModel) GetUserID() string {
	return u.UserID
}

// SetDepartmentID set the id of the user department
func (u *UserModel) SetDepartmentID(DepartmentID string) {
	u.DepartmentID = DepartmentID
}

// GetDepartmentID get the id of the user department
func (u *UserModel) GetDepartmentID() string {
	return u.DepartmentID
}

// SetUsername set the username of the user
func (u *UserModel) SetUsername(username string) {
	u.Username = username
}

// GetUsername get the username of the user
func (u *UserModel) GetUsername() string {
	return u.Username
}

// SetPassword set the user password
func (u *UserModel) SetPassword(password string) {
	u.Password = password
}

// GetPassword get the user password
func (u *UserModel) GetPassword() string {
	return u.Password
}

// SetName set the user name
func (u *UserModel) SetName(name string) {
	u.Name = name
}

// GetName get the user name
func (u *UserModel) GetName() string {
	return u.Name
}

// SetLastName set the user last name
func (u *UserModel) SetLastName(lastName string) {
	u.LastName = lastName
}

// GetLastName get the user last name
func (u *UserModel) GetLastName() string {
	return u.LastName
}

// SetEmail set the user email
func (u *UserModel) SetEmail(email string) {
	u.Email = email
}

// GetEmail get the user email
func (u *UserModel) GetEmail() string {
	return u.Email
}

// SetRole set the user roles
func (u *UserModel) SetRole(role []string) {
	u.Role = role
}

// GetRole get the user roles
func (u *UserModel) GetRole() []string {
	return u.Role
}

// SetActive set if the user is active
func (u *UserModel) SetActive(active bool) {
	u.Active = active
}

// GetActive get if the user is active
func (u *UserModel) GetActive() bool {
	return u.Active
}

// SetCreatedAt set the time when the user was created
func (u *UserModel) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

// GetCreatedAt get the time when the user was created
func (u *UserModel) GetCreatedAt() time.Time {
	return u.CreatedAt
}

// SetUpdatedAt set the time when the user was updated
func (u *UserModel) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

// GetUpdatedAt get the time when the user was updated
func (u *UserModel) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

// ToString Transform a UserModel to string
func (u *UserModel) ToString() string {
	return "\n***** UserModel Details *****\n" +
		"UserID= " + u.GetUserID() + "\n" +
		"DepartmentID= " + u.GetDepartmentID() + "\n" +
		"username= " + u.GetUsername() + "\n" +
		"password= " + u.GetPassword() + "\n" +
		"name= " + u.GetName() + "\n" +
		"lastName= " + u.GetLastName() + "\n" +
		"email= " + u.GetEmail() + "\n" +
		"role= " + strings.Join(u.GetRole()[:], ",") + "\n" +
		"active= " + strconv.FormatBool(u.GetActive()) + "\n" +
		"createdAt= " + u.GetCreatedAt().String() + "\n" +
		"updatedAt= " + u.GetUpdatedAt().String() + "\n" +
		"*******************************"
}
