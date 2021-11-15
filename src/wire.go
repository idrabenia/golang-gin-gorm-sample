//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	infgorm "example/hello/src/inf/gorm"
	"example/hello/src/model"
	"example/hello/src/services"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var set = wire.NewSet(
	services.ServiceSet,
	infgorm.RepoSet,
	provideDb,
	provideEngine,
)

type App struct {
	Engine      *gin.Engine
	UserService services.UserService
}

func InitApp() *App {
	wire.Build(set, wire.Struct(new(App), "*"))
	return &App{}
}

func provideDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.UserEntity{})

	return db
}

func provideEngine() *gin.Engine {
	return gin.Default()
}
