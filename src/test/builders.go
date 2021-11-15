package test

import (
	"example/hello/src/api"
	"example/hello/src/model"
	"gorm.io/gorm"
)

func MakeUser() *api.User {
	return &api.User{
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
