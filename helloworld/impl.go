package helloworld

import (
	"log"

	"github.com/AlexanderChen1989/recoverme/hello"
	"github.com/AlexanderChen1989/recoverme/world"
	"github.com/AlexanderChen1989/recoverme/world/bad"
	"github.com/AlexanderChen1989/recoverme/world/good"
)

// HelloWorlder implementation
type helloWorld struct {
	hello.Hellor
	world.Worlder
}

func (h *helloWorld) HelloWorld() error {
	log.Println(h.Hellor.Hello())
	log.Println(h.Worlder.World())
	return nil
}

// Create New HelloWorlder
func NewHelloGoodWorld() HelloWorlder {
	return &helloWorld{
		hello.AddRecover(hello.NewHello()),
		world.AddRecover(good.NewWorld()),
	}
}

// Create New HelloWorlder
func NewHelloBadWorld() HelloWorlder {
	return &helloWorld{
		hello.AddRecover(hello.NewHello()),
		world.AddRecover(bad.NewWorld()),
	}
}
