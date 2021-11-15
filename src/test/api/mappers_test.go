package api

import (
	"example/hello/src/api"
	"example/hello/src/model"
	"example/hello/src/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToUserEntity(t *testing.T) {
	user := test.MakeUser()

	result := api.ToUserEntity(user)

	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, result.LastName, "TestLastName")
}

func TestToUser(t *testing.T) {
	user := test.MakeUserEntity(1)

	result := api.ToUser(user)

	assert.Equal(t, result.Id, 1)
	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, result.LastName, "TestLastName")
}

func TestToUserList(t *testing.T) {
	firstUser := test.MakeUserEntity(1)
	secondUser := test.MakeUserEntity(2)
	users := []*model.UserEntity{firstUser, secondUser}

	result := api.ToUserList(users)

	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0].Id, 1)
	assert.Equal(t, result[1].Id, 2)
}
