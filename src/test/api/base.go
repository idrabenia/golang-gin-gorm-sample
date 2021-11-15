package api

import (
	model2 "example/hello/src/api/model"
	"example/hello/src/model"
	"gorm.io/gorm"
)

func MakeUser() *model2.User {
	return &model2.User{
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
