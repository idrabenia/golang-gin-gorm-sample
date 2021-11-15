package api

import (
	"example/hello/src/api/dto"
	"example/hello/src/model"
	"gorm.io/gorm"
)

func MakeUser() *dto.User {
	return &dto.User{
		Id:        1,
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
	}
}

func MakeUserEntity(id uint) *model.UserEntity {
	return &model.UserEntity{
		Model: gorm.Model{
			ID: id,
		},
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
	}
}
