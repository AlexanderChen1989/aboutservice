package world

import "fmt"

// Worlder Recover implementation
type worldRecover struct {
	Worlder
}

func (h *worldRecover) World() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("[panic][error] %s!", e)
		}
	}()
	return h.Worlder.World()
}

// Add Recover layer to Worlder
func AddRecover(h Worlder) Worlder {
	return &worldRecover{h}
}
