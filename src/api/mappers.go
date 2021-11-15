package api

import (
	"example/hello/src/api/dto"
	"example/hello/src/model"
)

func ToUserEntity(user *dto.User) *model.UserEntity {
	return &model.UserEntity{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUser(entity *model.UserEntity) *dto.User {
	return &dto.User{
		Id:        int(entity.ID),
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

func ToUserList(entities []*model.UserEntity) []*dto.User {
	result := make([]*dto.User, 0)

	for _, entity := range entities {
		result = append(result, ToUser(entity))
	}

	return result
}
