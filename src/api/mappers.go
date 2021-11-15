package api

import (
	"example/hello/src/model"
)

func ToUserEntity(user *model.User) *model.UserEntity {
	return &model.UserEntity{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUser(entity *model.UserEntity) *model.User {
	return &model.User{
		Id:        int(entity.ID),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

func ToUserList(entities []*model.UserEntity) []*model.User {
	result := make([]*model.User, 0)

	for _, entity := range entities {
		result = append(result, ToUser(entity))
	}

	return result
}
