package mapper

import (
	"Rein/domain"
	"Rein/presentation/model"
)

// UserModelDataMapper Mapper that transform a UserModel to a User
type UserModelDataMapper struct{}

// NewUserModelDataMapperr Constructor for the mapper
func NewUserModelDataMapperr() *UserModelDataMapper {
	u := new(UserModelDataMapper)
	return u
}

// Transform a UserModel to User
func (u *UserModelDataMapper) Transform(userModel model.UserModel) *domain.User {
	user := domain.NewUser()
	user.SetUserID(userModel.GetUserID())
	user.SetDepartmentID(userModel.GetDepartmentID())
	user.SetUsername(userModel.GetUsername())
	user.SetPassword(userModel.GetPassword())
	user.SetName(userModel.GetName())
	user.SetLastName(userModel.GetLastName())
	user.SetEmail(userModel.GetEmail())
	user.SetRole(userModel.GetRole())
	user.SetActive(userModel.GetActive())
	user.SetCreatedAt(userModel.GetCreatedAt())
	user.SetUpdatedAt(userModel.GetUpdatedAt())
	return user
}

// TransformSlice transform a Slice of UserEntity to Slice of User
func (u *UserModelDataMapper) TransformSlice(userModelSlice []model.UserModel) []*domain.User {
	userSlice := make([]*domain.User, len(userModelSlice))
	for i, userModel := range userModelSlice {
		userSlice[i] = u.Transform(userModel)
	}
	return userSlice
}

// Convert a User to UserEntity
func (u *UserModelDataMapper) Convert(user domain.User) *model.UserModel {
	userModel := model.NewUserModel()
	userModel.SetUserID(user.GetUserID())
	userModel.SetDepartmentID(user.GetDepartmentID())
	userModel.SetUsername(user.GetUsername())
	userModel.SetPassword(user.GetPassword())
	userModel.SetName(user.GetName())
	userModel.SetLastName(user.GetLastName())
	userModel.SetEmail(user.GetEmail())
	userModel.SetRole(user.GetRole())
	userModel.SetActive(user.GetActive())
	userModel.SetCreatedAt(user.GetCreatedAt())
	userModel.SetUpdatedAt(user.GetUpdatedAt())
	return userModel
}

// ConvertSlice transform a Slice of User to Slice of UserEntity
func (u *UserModelDataMapper) ConvertSlice(userSlice []domain.User) []*model.UserModel {
	userModelSlice := make([]*model.UserModel, len(userSlice))
	for i, user := range userSlice {
		userModelSlice[i] = u.Convert(user)
	}
	return userModelSlice
}
