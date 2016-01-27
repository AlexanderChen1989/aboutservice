package helloworld

import "fmt"

// HelloWorlder Recover implementation
type helloWorldRecover struct {
	HelloWorlder
}

func (h *helloWorldRecover) HelloWorld() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("[panic][error] %s!", e)
		}
	}()
	return h.HelloWorlder.HelloWorld()
}

// Add Recover layer to HelloWorlder
func AddRecover(h HelloWorlder) HelloWorlder {
	return &helloWorldRecover{h}
}
