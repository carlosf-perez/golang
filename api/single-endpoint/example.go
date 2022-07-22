package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func SetupEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

func HelloHandler(context *gin.Context) {
	helloResponse := HelloResponse{
		Message: "world",
	}

	context.JSON(http.StatusOK, helloResponse)
}

func main() {
	engine := SetupEngine()

	engine.GET("/hello", HelloHandler)

	engine.Run()
}
