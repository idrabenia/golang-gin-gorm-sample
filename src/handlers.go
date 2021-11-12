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
		id, err := ParseId(context)

		if err != nil {
			SendCode(context, 404)
			return
		}

		entity, err := userService.FindById(id)

		if err == nil {
			context.JSON(200, ToUser(entity))
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			SendCode(context, 404)
		} else {
			log.Println("Error on find user", id, err)
			SendCode(context, 500)
		}
	})

	r.GET("/user", func(context *gin.Context) {
		if users, err := userService.FindAll(); err == nil {
			context.JSON(200, ToUserList(users))
		} else {
			log.Println("Error on get all users " + err.Error())
			SendCode(context, 500)
		}
	})

	r.POST("/user", func(context *gin.Context) {
		user := User{}

		if err := context.ShouldBind(&user); err != nil {
			log.Println(err)
			SendCode(context, 400)
			return
		}

		entity, err := userService.Create(ToUserEntity(&user))

		if err == nil {
			context.JSON(200, ToUser(entity))
		} else {
			log.Println(err)
			SendCode(context, 500)
		}
	})

	r.PUT("/user/:id", func(context *gin.Context) {
		id, err := ParseId(context)

		if err != nil {
			SendCode(context, 404)
			return
		}

		command := UpdateUserCommand{}

		if err := context.ShouldBind(&command); err != nil {
			log.Println(err)
			SendCode(context, 400)
			return
		}

		if entity, err := userService.Update(id, &command); err == nil {
			context.JSON(200, ToUser(entity))
		} else {
			log.Println("Error on update user ", err.Error())
			SendCode(context, 500)
		}
	})

	r.DELETE("/user/:id", func(context *gin.Context) {
		id, err := ParseId(context)

		if err != nil {
			SendCode(context, 404)
			return
		}

		if result := userService.Delete(id); result == nil {
			SendCode(context, 200)
		} else {
			SendCode(context, 500)
		}
	})

}

func SendCode(context *gin.Context, code int) {
	context.Writer.WriteHeader(code)
}

func ParseId(context *gin.Context) (int, error) {
	id := context.Params.ByName("id")
	idVal, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Could not parser user ID", err)
		return idVal, err
	} else {
		return idVal, nil
	}
}

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserEntity{})

	return db
}
