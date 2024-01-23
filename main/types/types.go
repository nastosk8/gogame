package types

type OutputMessage struct {
	Message string
	PlayerCounter int
}

type Class struct {
	Name       string
	DamageType string
	Health     int
	Mana       int
}

type Hero struct {
	Name  string
	Class Class
}