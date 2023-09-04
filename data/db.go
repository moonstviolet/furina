package data

import (
	"encoding/json"
	"errors"
	"furina/storage"
	"strings"
)

type DB struct {
	storage.Storage
}

const (
	UserProfileKey = "profile"
	IdSeparator    = "_"
)

var (
	gDb *DB
)

func getDB() *DB {
	return gDb
}

func getPrefixAndKey(id string, v any) (pre, key string, err error) {
	switch v.(type) {
	case User, *User:
		pre, key = id, UserProfileKey
	case Character, *Character:
		pre, key, err = getPrefixAndKeyById(id)
	default:
		err = errors.New("错误的数据对象")
	}
	return
}

func (db *DB) Get(id string, v any) error {
	prefix, key, err := getPrefixAndKey(id, v)
	if err != nil {
		return err
	}
	bytes, err := db.Storage.Read(prefix, key)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		return nil
	}
	return json.Unmarshal(bytes, v)
}

func (db *DB) Put(id string, v any) error {
	prefix, key, err := getPrefixAndKey(id, v)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return db.Storage.Write(prefix, key, bytes)
}

func getIdByPrefixAndKey(pre, key string) string {
	return pre + IdSeparator + key
}

func getPrefixAndKeyById(id string) (pre, key string, err error) {
	list := strings.Split(id, IdSeparator)
	if len(list) != 2 {
		err = errors.New("错误的ID")
		return
	}
	return list[0], list[1], nil
}

func init() {
	gDb = &DB{
		Storage: storage.NewLocalStorage(),
	}
}
