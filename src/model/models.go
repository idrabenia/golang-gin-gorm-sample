package model

import "gorm.io/gorm"

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type UserEntity struct {
	gorm.Model
	FirstName string
	LastName  string
}

type UpdateUserCommand struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type GormDb interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}
