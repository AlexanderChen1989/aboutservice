package hello

import "fmt"

// Hellor Recover implementation
type helloRecover struct {
	Hellor
}

func (h *helloRecover) Hello() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("[panic][error] %s!", e)
		}
	}()
	return h.Hellor.Hello()
}

// Add Recover layer to Hellor
func AddRecover(h Hellor) Hellor {
	return &helloRecover{h}
}
