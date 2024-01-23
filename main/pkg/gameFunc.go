package gameFunc

import (
	"Game/types"
	"database/sql"
	"fmt"
	"time"
)

var db, err = sql.Open("mysql", "root:AdmPass1234@/gogame")
var PlayerCounter = 0

func main(){
	fmt.Println("Hello")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func Game(AttackingPlayer types.Hero, DeffendingPlayer types.Hero)types.OutputMessage {
	
	if PlayerCounter == 0 {
		PlayerCounter++
	} else {
		PlayerCounter--
	}

	return types.OutputMessage{
		Message: "Attacking player is " + AttackingPlayer.Name + " Defending player is " + DeffendingPlayer.Name + "." ,
		PlayerCounter: PlayerCounter ,
	}
}

