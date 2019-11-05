package internal

type Repositories struct {
	Model interface{}
}

// Repositories responsible
// to handle data input and database
type RepositoriesContract interface {
	Select()
}

func (r Repositories) Select() {
}
