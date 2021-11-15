package test

import (
	"example/hello/src/model"
	"gorm.io/gorm"
)

func MakeUser() *model.User {
	return &model.User{
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
