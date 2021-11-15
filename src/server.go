package main

import (
	"example/hello/src/api"
)

func main() {
	app := InitApp()

	api.Handlers(app.Engine, app.UserService)

	app.Engine.Run(":8080")
}
