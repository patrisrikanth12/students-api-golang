package storage

type Storage interface {
	CreateStudent(name string, email string, mobile string) (int64, error)
}