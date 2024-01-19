package types

type Class struct {
	name string
	damage_type string
	health int
	mana int
}

type Hero struct {
	name string
	class Class
}