package storage

type Storage interface {
	Start() error
	Write(table, id string, value any) error
	Read(table, id string, value any) error
}
