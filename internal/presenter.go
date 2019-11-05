package internal

// Presenter struct used with any entities
// To present status code, error and it data
type Presenter struct {
	StatusCode int
	Message interface{}
	Data interface{}
}

