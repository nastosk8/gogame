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
	Oppositeplayer := 1
	if gameFunc.PlayerCounter == 0 {
		Oppositeplayer = 1
	}else {
		Oppositeplayer = 0
	}
	result := gameFunc.Game(data.AllPlayers[gameFunc.PlayerCounter], data.AllPlayers[Oppositeplayer])
	webServerContext.JSON(http.StatusOK, gin.H{
		"message": result.Message,
		"playerCounter": result.PlayerCounter,
		"OppositePlayerCounter": Oppositeplayer,
	  })
}
