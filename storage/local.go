package storage

import (
	"furina/logger"
	"os"
	"path/filepath"

	"furina/config"
)

type LocalStorage struct {
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) Start() error {
	return nil
}

func (s *LocalStorage) Write(prefix, key string, data []byte) error {
	createIfNotExist(prefix, key)
	path := filepath.Join(config.GetLocalStorageDir(), prefix, key)
	return os.WriteFile(path, data, 0666)
}

func (s *LocalStorage) Read(prefix, key string) (data []byte, err error) {
	createIfNotExist(prefix, key)
	path := filepath.Join(config.GetLocalStorageDir(), prefix, key)
	data, err = os.ReadFile(path)
	return
}

func createIfNotExist(prefix, key string) {
	dir := filepath.Join(config.GetLocalStorageDir(), prefix)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0744)
		logger.Error("创建目录失败", "err", err)
	}
	path := filepath.Join(dir, key)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		logger.Error("创建文件失败", "err", err)
	}
}
