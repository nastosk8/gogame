package gameFunc

import (
	"Game/types"
	"database/sql"
	"fmt"
	"time"
)

var db, err = sql.Open("mysql", "root:AdmPass1234@/gogame")
var CurrentPlayer = 0
var OppositePlayer = 1

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
	if CurrentPlayer == 0 {
		CurrentPlayer++
		OppositePlayer--
	} else {
		CurrentPlayer--
		OppositePlayer++
	}

	return types.OutputMessage{
		Message: "Attacking player is " + AttackingPlayer.Name + " Defending player is " + DeffendingPlayer.Name + "." ,
		CurrentPlayer: CurrentPlayer ,
		OppositePlayer: OppositePlayer ,
	}
}


