package gorm

import (
	"example/hello/src/model"
	"example/hello/src/test/api"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) FindById(id int) (*model.UserEntity, error) {
	m.Called(id)
	return api.MakeUserEntity(1), nil
}

func (m *UserRepoMock) Create(entity *model.UserEntity) (*model.UserEntity, error) {
	m.Called(entity)
	return api.MakeUserEntity(1), nil
}

func (m *UserRepoMock) FindAll() ([]*model.UserEntity, error) {
	m.Called()
	return []*model.UserEntity{api.MakeUserEntity(1)}, nil
}

func (m *UserRepoMock) Update(command *model.UserEntity) (*model.UserEntity, error) {
	m.Called(command)
	return api.MakeUserEntity(1), nil
}

func (m *UserRepoMock) Delete(id int) error {
	m.Called(id)
	return nil
}
