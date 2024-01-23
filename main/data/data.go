package data

import "Game/types"

var AllClasses map[string]types.Class = map[string]types.Class{
	"Mage": {Name: "Mage", DamageType: "Magical", Health: 80, Mana: 100},
	"Warrior": {Name: "Warrior", DamageType: "Psysical", Health: 100, Mana: 0},
}
var AllPlayers []types.Hero = []types.Hero{
	{Name: "Jakub1", Class: AllClasses["Warrior"]},
	{Name: "Jakub2", Class: AllClasses["Mage"]},
}


