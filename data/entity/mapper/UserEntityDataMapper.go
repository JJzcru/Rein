package mapper

import (
	"Rein/data/entity"
	"Rein/domain"
)

// UserEntityDataMapper Mapper that transform a UserEntity to a User
type UserEntityDataMapper struct{}

// NewUserEntityDataMapper Constructor for the mapper
func NewUserEntityDataMapper() *UserEntityDataMapper {
	u := new(UserEntityDataMapper)
	return u
}

// Transform a UserEntity to User
func (u *UserEntityDataMapper) Transform(userEntity entity.UserEntity) *domain.User {
	user := domain.NewUser()
	user.SetUserID(userEntity.GetUserID())
	user.SetDepartmentID(userEntity.GetDepartmentID())
	user.SetUsername(userEntity.GetUsername())
	user.SetPassword(userEntity.GetPassword())
	user.SetName(userEntity.GetName())
	user.SetLastName(userEntity.GetLastName())
	user.SetEmail(userEntity.GetEmail())
	user.SetRole(userEntity.GetRole())
	user.SetActive(userEntity.GetActive())
	user.SetCreatedAt(userEntity.GetCreatedAt())
	user.SetUpdatedAt(userEntity.GetUpdatedAt())
	return user
}

// TransformSlice transform a Slice of UserEntity to Slice of User
func (u *UserEntityDataMapper) TransformSlice(userEntitySlice []entity.UserEntity) []*domain.User {
	userSlice := make([]*domain.User, len(userEntitySlice))
	for i, userEntity := range userEntitySlice {
		userSlice[i] = u.Transform(userEntity)
	}
	return userSlice
}

// Convert a User to UserEntity
func (u *UserEntityDataMapper) Convert(user domain.User) *entity.UserEntity {
	userEntity := entity.NewUserEntity()
	userEntity.SetUserID(user.GetUserID())
	userEntity.SetDepartmentID(user.GetDepartmentID())
	userEntity.SetUsername(user.GetUsername())
	userEntity.SetPassword(user.GetPassword())
	userEntity.SetName(user.GetName())
	userEntity.SetLastName(user.GetLastName())
	userEntity.SetEmail(user.GetEmail())
	userEntity.SetRole(user.GetRole())
	userEntity.SetActive(user.GetActive())
	userEntity.SetCreatedAt(user.GetCreatedAt())
	userEntity.SetUpdatedAt(user.GetUpdatedAt())
	return userEntity
}

// ConvertSlice transform a Slice of User to Slice of UserEntity
func (u *UserEntityDataMapper) ConvertSlice(userSlice []domain.User) []*entity.UserEntity {
	userEntitySlice := make([]*entity.UserEntity, len(userSlice))
	for i, user := range userSlice {
		userEntitySlice[i] = u.Convert(user)
	}
	return userEntitySlice
}
