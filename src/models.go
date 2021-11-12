package main

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
