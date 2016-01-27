package main

import (
	"log"

	"github.com/AlexanderChen1989/aboutservice/helloworld"
)

func main() {
	log.SetFlags(log.Llongfile)
	h := helloworld.NewHelloGoodWorld()
	h = helloworld.AddRecover(h)
	log.Println(h.HelloWorld())

	h = helloworld.NewHelloBadWorld()
	h = helloworld.AddRecover(h)
	log.Println(h.HelloWorld())
}
