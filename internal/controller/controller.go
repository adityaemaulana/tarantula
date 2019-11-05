package controller

import "github.com/jeremylombogia/tarantula/internal"

// Controller model that filled
// With Payload and Presenter
type Model struct {
	Payload   internal.Payload
	Presenter internal.Presenter
}

// ControllerHandler contract
type Controller interface {
	Get() internal.Presenter
	Post() internal.Presenter
	Put() internal.Presenter
	Delete() internal.Presenter
}

// Get interface controller
func (c Model) Get() internal.Presenter {
	return c.Presenter
}

// Post interface controller
func (c Model) Post() internal.Presenter {
	return c.Presenter
}

// Put interface controller
func (c Model) Put() internal.Presenter {
	return c.Presenter
}

// Delete interface controller
func (c Model) Delete() internal.Presenter {
	return c.Presenter
}
