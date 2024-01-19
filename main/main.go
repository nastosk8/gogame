package main

import (
	"fmt"
	"net/http"
	"github.com/go-sql-driver/mysql"
	"Game/types"
	"Game/data"
	"github.com/gin-gonic/gin"
)

func game(AttackingPlayer types.Hero, DeffendingPlayer types.Hero){
	fmt.Print("Attacking player is %s", AttackingPlayer)
	fmt.Print("Defending player is %s", DeffendingPlayer)
}

func main() {
	WebServer := gin.Default()
	WebServer.POST("/play", CallHanlder)
}

func callHanlder(webServerContext *gin.Context) {
	
}
