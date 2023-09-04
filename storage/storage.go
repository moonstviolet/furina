package storage

type Storage interface {
	Start() error
	Write(prefix, key string, data []byte) error
	Read(prefix, key string) (data []byte, err error)
}
