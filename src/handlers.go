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

	r.GET("/user", func(context *gin.Context) {
		if users, err := userService.FindAll(); err == nil {
			context.JSON(200, ToUserList(users))
		} else {
			log.Println("Error on get all users " + err.Error())
			context.Writer.WriteHeader(500)
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

	r.PUT("/user/:id", func(context *gin.Context) {
		id := context.Params.ByName("id")
		idVal, err := strconv.Atoi(id)

		if err != nil {
			log.Println(err)
			context.Writer.WriteHeader(404)
			return
		}

		command := UpdateUserCommand{}

		if err := context.ShouldBind(&command); err != nil {
			log.Println(err)
			context.Writer.WriteHeader(400)
			return
		}

		if entity, err := userService.Update(idVal, &command); err == nil {
			context.JSON(200, ToUser(entity))
		} else {
			log.Println("Error on update user ", err.Error())
			context.Writer.WriteHeader(500)
		}
	})

	r.DELETE("/user/:id", func(context *gin.Context) {
		id := context.Params.ByName("id")
		idVal, err := strconv.Atoi(id)

		if err != nil {
			log.Println(err)
			context.Writer.WriteHeader(404)
			return
		}

		if result := userService.Delete(idVal); result == nil {
			context.Writer.WriteHeader(200)
		} else {
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
