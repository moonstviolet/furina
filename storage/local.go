package storage

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"furina/config"
)

type LocalStorage struct {
}

func NewLocalStorage() Storage {
	return &LocalStorage{}
}

func (s *LocalStorage) Start() error {
	return nil
}

func (s *LocalStorage) Write(table, id string, value any) error {
	path := createIfNotExist(table, id)
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0666)
}

func (s *LocalStorage) Read(table, id string, value any) error {
	path := createIfNotExist(table, id)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, value)
}

func createIfNotExist(table, id string) string {
	dir := filepath.Join(config.GetLocalStorageDir(), table)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0744)
		if err != nil {
			log.Fatalln("创建目录失败, err:", err)
		}
	}
	path := filepath.Join(dir, id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			log.Fatalln("创建文件失败, err:", err)
		}
	}
	return path
}
