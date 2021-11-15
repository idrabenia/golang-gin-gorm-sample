package api

import (
	model2 "example/hello/src/api/model"
	"example/hello/src/model"
)

func ToUserEntity(user *model2.User) *model.UserEntity {
	return &model.UserEntity{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUser(entity *model.UserEntity) *model2.User {
	return &model2.User{
		Id:        int(entity.ID),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

func ToUserList(entities []*model.UserEntity) []*model2.User {
	result := make([]*model2.User, 0)

	for _, entity := range entities {
		result = append(result, ToUser(entity))
	}

	return result
}
