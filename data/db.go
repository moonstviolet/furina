package data

import (
	"errors"
	"furina/config"
	"furina/storage"
	"log"
	"strings"
)

type DB struct {
	storage.Storage
}

const (
	TableNameUser      = "user"
	TableNameCharacter = "character"
	IdSeparator        = "_"
)

var (
	gDb = &DB{}
)

func getDB() *DB {
	return gDb
}

func (db *DB) Get(table, id string, v any) error {
	return db.Storage.Read(table, id, v)
}

func (db *DB) Put(table, id string, v any) error {
	return db.Storage.Write(table, id, v)
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
	dbConfig := config.GetConfig().Database
	if dbConfig.Addr == "" {
		gDb.Storage = storage.NewLocalStorage()
	} else {
		gDb.Storage = storage.NewMongoDB(
			storage.NewMongoDBReq{
				DBName: dbConfig.DBName,
				Addr:   dbConfig.Addr,
			},
		)
	}
	err := gDb.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
