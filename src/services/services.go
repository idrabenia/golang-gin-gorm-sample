package services

import (
	"example/hello/src/model"
)

type UserService interface {
	FindById(id int) (*model.UserEntity, error)

	Create(entity *model.UserEntity) (*model.UserEntity, error)

	FindAll() ([]*model.UserEntity, error)

	Update(id int, command *UpdateUserCommand) (*model.UserEntity, error)

	Delete(id int) error
}

type UserServiceImpl struct {
	UserRepo model.UserRepo
}

func (s *UserServiceImpl) FindById(id int) (*model.UserEntity, error) {
	return s.UserRepo.FindById(id)
}

func (s *UserServiceImpl) Create(entity *model.UserEntity) (*model.UserEntity, error) {
	return s.UserRepo.Create(entity)
}

func (s *UserServiceImpl) FindAll() ([]*model.UserEntity, error) {
	return s.UserRepo.FindAll()
}

func (s *UserServiceImpl) Update(id int,
	command *UpdateUserCommand) (*model.UserEntity, error) {

	if entity, err := s.UserRepo.FindById(id); err == nil {
		entity.FirstName = command.FirstName
		entity.LastName = command.LastName

		return s.UserRepo.Update(entity)
	} else {
		return nil, err
	}
}

func (s *UserServiceImpl) Delete(id int) error {
	return s.UserRepo.Delete(id)
}
