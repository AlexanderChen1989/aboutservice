package good

import (
	"log"

	"github.com/AlexanderChen1989/recoverme/world"
)

// Worlder implementation
type goodWorld struct{}

func (h *goodWorld) World() error {
	log.Println("Good World Not Panic!")
	return nil
}

// Create New Worlder
func NewWorld() world.Worlder {
	return &goodWorld{}
}
