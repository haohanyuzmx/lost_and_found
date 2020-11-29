package main

import (
	"github.com/gin-gonic/gin"
	"lost_and_found/api"
)

func main() {
	engine := gin.Default()
	engine.Static("picture", "")
	api.User(engine)
	api.Post(engine)
	engine.Run(":8080")
}
