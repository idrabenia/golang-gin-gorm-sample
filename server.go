package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := InitDb()

	Handlers(db, r)

	r.Run(":8080")
}
