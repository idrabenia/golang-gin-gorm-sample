package model

import (
	"example/hello/src/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DbMock struct {
	mock.Mock
}

func (db *DbMock) Create(value interface{}) (tx *gorm.DB) {
	db.Called(value)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) Updates(values interface{}) (tx *gorm.DB) {
	db.Called(values)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.Called(value, conds)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.Called(dest, conds)
	return &gorm.DB{Error: nil}
}

func (db *DbMock) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	db.Called(dest, conds)

	item := dest.(*model.UserEntity)
	item.ID = 1

	return &gorm.DB{Error: nil}
}
