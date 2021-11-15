package services

import (
	"example/hello/src/model"
	"example/hello/src/test/api"
	mockdb "example/hello/src/test/inf/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFindById(t *testing.T) {
	repo := mockDb(new(mockdb.UserRepoMock))
	service := UserServiceImpl{UserRepo: repo}

	result, err := service.FindById(1)

	assert.Equal(t, result.ID, uint(1))
	assert.Equal(t, err, nil)
}

func TestCreate(t *testing.T) {
	repo := mockDb(new(mockdb.UserRepoMock))
	service := UserServiceImpl{UserRepo: repo}
	entity := api.MakeUserEntity(1)

	result, err := service.Create(entity)

	assert.Equal(t, result.ID, uint(1))
	assert.Equal(t, err, nil)
}

func TestFindAll(t *testing.T) {
	repo := mockDb(new(mockdb.UserRepoMock))
	service := UserServiceImpl{UserRepo: repo}

	result, err := service.FindAll()

	assert.Equal(t, len(result), 1)
	assert.Equal(t, err, nil)
}

func TestDelete(t *testing.T) {
	repo := mockDb(new(mockdb.UserRepoMock))
	service := UserServiceImpl{UserRepo: repo}

	err := service.Delete(1)

	assert.Equal(t, err, nil)
	repo.AssertCalled(t, "Delete", mock.Anything, mock.Anything)
}

func TestUpdate(t *testing.T) {
	repo := mockDb(new(mockdb.UserRepoMock))
	service := UserServiceImpl{UserRepo: repo}
	command := MakeUpdateCommand()

	result, err := service.Update(1, command)

	assert.Equal(t, result.FirstName, "TestFirstName")
	assert.Equal(t, err, nil)
}

func mockDb(repoMock *mockdb.UserRepoMock) *mockdb.UserRepoMock {
	repoMock.
		On("FindAll").
		Return([]*model.UserEntity{api.MakeUserEntity(1)})

	repoMock.
		On("FindById", 1).
		Return(api.MakeUserEntity(1), nil)

	repoMock.
		On("Delete", 1).
		Return(api.MakeUserEntity(1), nil)

	repoMock.
		On("Create", mock.Anything).
		Return(api.MakeUserEntity(1), nil)

	repoMock.
		On("Update", mock.Anything).
		Return(api.MakeUserEntity(1), nil)

	return repoMock
}

func MakeUpdateCommand() *UpdateUserCommand {
	return &UpdateUserCommand{
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
	}
}
