package model

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	FirstName string
	LastName  string
}
