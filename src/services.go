package main

import (
	"gorm.io/gorm"
)

type GormDb interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

type UserServiceType interface {
	FindById(id int) (*UserEntity, error)
	Create(entity *UserEntity) (*UserEntity, error)
	FindAll() ([]*UserEntity, error)
	Update(id int, command *UpdateUserCommand) (*UserEntity, error)
	Delete(id int) error
}

type UserService struct {
	Db GormDb
}

func (service *UserService) FindById(id int) (*UserEntity, error) {
	entity := UserEntity{}
	result := service.Db.First(&entity, id)

	if result.Error == nil {
		return &entity, nil
	} else {
		return nil, result.Error
	}
}

func (service *UserService) Create(entity *UserEntity) (*UserEntity, error) {
	result := service.Db.Create(&entity)

	if result.Error == nil {
		return entity, nil
	} else {
		return nil, result.Error
	}
}

func (service *UserService) FindAll() ([]*UserEntity, error) {
	var users []*UserEntity

	result := service.Db.Find(&users)

	if result.Error == nil {
		return users, nil
	} else {
		return nil, result.Error
	}
}

func (service *UserService) Update(id int, command *UpdateUserCommand) (*UserEntity, error) {
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
