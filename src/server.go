package main

import (
	"example/hello/src/api"
	"example/hello/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	db := InitDb()

	api.Handlers(db, r)

	r.Run(":8080")
}

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.UserEntity{})

	return db
}
