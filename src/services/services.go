package services

import (
	"example/hello/src/model"
)

type UserServiceType interface {
	FindById(id int) (*model.UserEntity, error)
	Create(entity *model.UserEntity) (*model.UserEntity, error)
	FindAll() ([]*model.UserEntity, error)
	Update(id int, command *model.UpdateUserCommand) (*model.UserEntity, error)
	Delete(id int) error
}

type UserService struct {
	Db model.GormDb
}

func (service *UserService) FindById(id int) (*model.UserEntity, error) {
	entity := model.UserEntity{}
	result := service.Db.First(&entity, id)

	if result.Error == nil {
		return &entity, nil
	} else {
		return nil, result.Error
	}
}

func (service *UserService) Create(entity *model.UserEntity) (*model.UserEntity, error) {
	result := service.Db.Create(&entity)

	if result.Error == nil {
		return entity, nil
	} else {
		return nil, result.Error
	}
}

func (service *UserService) FindAll() ([]*model.UserEntity, error) {
	var users []*model.UserEntity

	result := service.Db.Find(&users)

	if result.Error == nil {
		return users, nil
	} else {
		return nil, result.Error
	}
}

func (service *UserService) Update(id int, command *model.UpdateUserCommand) (*model.UserEntity, error) {
	if entity, err := service.FindById(id); err == nil {
		entity.FirstName = command.FirstName
		entity.LastName = command.LastName

		if result := service.Db.Updates(entity); result.Error == nil {
			return entity, nil
		} else {
			return nil, result.Error
		}
	} else {
		return nil, err
	}
}

func (service *UserService) Delete(id int) error {
	user, err := service.FindById(id)

	if err != nil {
		return err
	}

	if res := service.Db.Delete(user); res.Error != nil {
		return res.Error
	}

	return nil
}
