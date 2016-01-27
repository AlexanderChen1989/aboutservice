package bad

import "github.com/AlexanderChen1989/recoverme/world"

// Worlder implementation
type badWorld struct{}

func (h *badWorld) World() error {
	panic("Bad World Panic")
}

// Create New Worlder
func NewWorld() world.Worlder {
	return &badWorld{}
}
