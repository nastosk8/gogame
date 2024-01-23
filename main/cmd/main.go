package main

import (
	"Game/data"
	gameFunc "Game/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	webServer := gin.Default()
	webServer.POST("/play", callHanlder)
	webServer.Run()
}

func callHanlder(webServerContext *gin.Context) {
	result := gameFunc.Game(data.AllPlayers[gameFunc.CurrentPlayer], data.AllPlayers[gameFunc.OppositePlayer])
	webServerContext.JSON(http.StatusOK, gin.H{
		"Message": result.Message,
		"PlayerCounter": result.CurrentPlayer,
		"OppositePlayer": result.OppositePlayer,
	  })
}
