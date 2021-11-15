package services

import (
	"example/hello/src/model"
	"example/hello/src/services"
	"example/hello/src/test"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) FindById(id int) (*model.UserEntity, error) {
	m.Called(id)
	return test.MakeUserEntity(1), nil
}

func (m *UserServiceMock) Create(entity *model.UserEntity) (*model.UserEntity, error) {
	m.Called(entity)
	return test.MakeUserEntity(1), nil
}

func (m *UserServiceMock) FindAll() ([]*model.UserEntity, error) {
	m.Called()
	return []*model.UserEntity{test.MakeUserEntity(1)}, nil
}

func (m *UserServiceMock) Update(id int,
	command *services.UpdateUserCommand) (*model.UserEntity, error) {

	m.Called(id, command)
	return test.MakeUserEntity(1), nil
}

func (m *UserServiceMock) Delete(id int) error {
	m.Called(id)
	return nil
}
