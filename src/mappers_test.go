package main

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestToUserEntity(t *testing.T) {
	user := MakeUser()

	result := ToUserEntity(user)

	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, result.LastName, "TestLastName")
}

func TestToUser(t *testing.T) {
	user := MakeUserEntity(1)

	result := ToUser(user)

	assert.Equal(t, result.Id, 1)
	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, result.LastName, "TestLastName")
}

func TestToUserList(t *testing.T) {
	firstUser := MakeUserEntity(1)
	secondUser := MakeUserEntity(2)
	users := []*UserEntity{firstUser, secondUser}

	result := ToUserList(users)

	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0].Id, 1)
	assert.Equal(t, result[1].Id, 2)
}

func MakeUser() *User {
	return &User{
		Id:        1,
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
	}
}

func MakeUserEntity(id uint) *UserEntity {
	return &UserEntity{
		Model: gorm.Model{
			ID: id,
		},
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
	}
}
