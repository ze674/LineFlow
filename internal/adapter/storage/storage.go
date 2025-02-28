package storage

const (
	ErrNotFound = "not found"
	ErrExists   = "exists"
)

type Storage interface {
	Save(data string) error
}
