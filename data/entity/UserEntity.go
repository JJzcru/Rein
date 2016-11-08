package entity

import (
	"strconv"
	"strings"
	"time"
)

// UserEntity Representation of User in the data layer
type UserEntity struct {
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

// NewUserEntity Constructor for a new UserEntity
func NewUserEntity() *UserEntity {
	return &UserEntity{}
}

// SetUserID set the id of the user
func (u *UserEntity) SetUserID(userID string) {
	u.userID = userID
}

// GetUserID get the id of the user
func (u *UserEntity) GetUserID() string {
	return u.userID
}

// SetDepartmentID set the id of the user department
func (u *UserEntity) SetDepartmentID(departmentID string) {
	u.departmentID = departmentID
}

// GetDepartmentID get the id of the user department
func (u *UserEntity) GetDepartmentID() string {
	return u.departmentID
}

// SetUsername set the username of the user
func (u *UserEntity) SetUsername(username string) {
	u.username = username
}

// GetUsername get the username of the user
func (u *UserEntity) GetUsername() string {
	return u.username
}

// SetPassword set the user password
func (u *UserEntity) SetPassword(password string) {
	u.password = password
}

// GetPassword get the user password
func (u *UserEntity) GetPassword() string {
	return u.password
}

// SetName set the user name
func (u *UserEntity) SetName(name string) {
	u.name = name
}

// GetName get the user name
func (u *UserEntity) GetName() string {
	return u.name
}

// SetLastName set the user last name
func (u *UserEntity) SetLastName(lastName string) {
	u.lastName = lastName
}

// GetLastName get the user last name
func (u *UserEntity) GetLastName() string {
	return u.lastName
}

// SetEmail set the user email
func (u *UserEntity) SetEmail(email string) {
	u.email = email
}

// GetEmail get the user email
func (u *UserEntity) GetEmail() string {
	return u.email
}

// SetRole set the user roles
func (u *UserEntity) SetRole(role []string) {
	u.role = role
}

// GetRole get the user roles
func (u *UserEntity) GetRole() []string {
	return u.role
}

// SetActive set if the user is active
func (u *UserEntity) SetActive(active bool) {
	u.active = active
}

// GetActive get if the user is active
func (u *UserEntity) GetActive() bool {
	return u.active
}

// SetCreatedAt set the time when the user was created
func (u *UserEntity) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

// GetCreatedAt get the time when the user was created
func (u *UserEntity) GetCreatedAt() time.Time {
	return u.createdAt
}

// SetUpdatedAt set the time when the user was updated
func (u *UserEntity) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

// GetUpdatedAt get the time when the user was updated
func (u *UserEntity) GetUpdatedAt() time.Time {
	return u.updatedAt
}

// ToString Transform a UserEntity to string
func (u *UserEntity) ToString() string {
	return "\n***** UserEntity Details *****\n" +
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
