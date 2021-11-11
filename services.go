package main

import (
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
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
