package hello

// Hellor implementation
type hello struct{}

func (h *hello) Hello() error {
	panic("Hello Panic")
}

// Create New Hellor
func NewHello() Hellor {
	return &hello{}
}
