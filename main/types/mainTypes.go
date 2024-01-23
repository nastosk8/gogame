package types

type RoundInterface interface {
	TakeDamage() DamageEvaluation
}

type DamageEvaluation struct {
	BeforeHealth int
	AfterHealth  int
	DamageType   string
}

type OutputMessage struct {
	Message        string
	CurrentPlayer  int
	OppositePlayer int
}

type SpellClass struct {
	Name       string
	DamageType string
	Health     int
	Mana       int
}

type PhysicalClass struct {
	Name       string
	DamageType string
	Health     int
	Stamina    int
}

type Hero struct {
	Name string
}