package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func Handlers(db *gorm.DB, r *gin.Engine) {

	userService := UserService{Db: db}

	r.GET("/user/:id", func(context *gin.Context) {
		id := context.Params.ByName("id")
		idVal, err := strconv.Atoi(id)

		if err != nil {
			log.Println(err)
			context.Writer.WriteHeader(404)
			return
		}

		entity, err := userService.FindById(idVal)

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(200, ToUser(entity))
		} else {
			context.Writer.WriteHeader(404)
		}
	})

	r.POST("/user", func(context *gin.Context) {
		user := User{}

		if err := context.ShouldBind(&user); err != nil {
			log.Println(err)
			context.Writer.WriteHeader(400)
			return
		}

		entity, err := userService.Create(ToUserEntity(&user))

		if err == nil {
			context.JSON(200, ToUser(entity))
		} else {
			log.Println(err)
			context.Writer.WriteHeader(500)
		}
	})

}

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserEntity{})

	return db
}
