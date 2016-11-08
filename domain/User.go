package domain

import (
	"strconv"
	"strings"
	"time"
)

// User Representation of an User
type User struct {
	userID       string
	departmentID string
	username     string
	password     string
	name         string
	lastName     string
	email        string
	role         []string
	active       bool
	createdAt    time.Time
	updatedAt    time.Time
}

// NewUser Constructor for a new User
func NewUser() *User {
	return &User{}
}

// SetUserID set the id of the user
func (u *User) SetUserID(userID string) {
	u.userID = userID
}

// GetUserID get the id of the user
func (u *User) GetUserID() string {
	return u.userID
}

// SetDepartmentID set the id of the user department
func (u *User) SetDepartmentID(departmentID string) {
	u.departmentID = departmentID
}

// GetDepartmentID get the id of the user department
func (u *User) GetDepartmentID() string {
	return u.departmentID
}

// SetUsername set the username of the user
func (u *User) SetUsername(username string) {
	u.username = username
}

// GetUsername get the username of the user
func (u *User) GetUsername() string {
	return u.username
}

// SetPassword set the user password
func (u *User) SetPassword(password string) {
	u.password = password
}

// GetPassword get the user password
func (u *User) GetPassword() string {
	return u.password
}

// SetName set the user name
func (u *User) SetName(name string) {
	u.name = name
}

// GetName get the user name
func (u *User) GetName() string {
	return u.name
}

// SetLastName set the user last name
func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

// GetLastName get the user last name
func (u *User) GetLastName() string {
	return u.lastName
}

// SetEmail set the user email
func (u *User) SetEmail(email string) {
	u.email = email
}

// GetEmail get the user email
func (u *User) GetEmail() string {
	return u.email
}

// SetRole set the user roles
func (u *User) SetRole(role []string) {
	u.role = role
}

// GetRole get the user roles
func (u *User) GetRole() []string {
	return u.role
}

// SetActive set if the user is active
func (u *User) SetActive(active bool) {
	u.active = active
}

// GetActive get if the user is active
func (u *User) GetActive() bool {
	return u.active
}

// SetCreatedAt set the time when the user was created
func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

// GetCreatedAt get the time when the user was created
func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

// SetUpdatedAt set the time when the user was updated
func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

// GetUpdatedAt get the time when the user was updated
func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

// ToString Transform a User to string
func (u *User) ToString() string {
	return "\n***** User Details *****\n" +
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
