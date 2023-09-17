package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DBName string
	Addr   string
	Db     *mongo.Database
	Client *mongo.Client
}

type NewMongoDBReq struct {
	Addr   string
	DBName string
}

func NewMongoDB(req NewMongoDBReq) Storage {
	return &MongoDB{
		DBName: req.DBName,
		Addr:   req.Addr,
	}
}

func (m *MongoDB) Start() error {
	uri := fmt.Sprintf("mongodb://%s/?connect=direct", m.Addr)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	m.Client = client
	m.Db = client.Database(m.DBName)
	return nil
}

func (m *MongoDB) Write(table, id string, value any) error {
	_, err := m.Db.Collection(table).ReplaceOne(
		context.Background(), bson.M{"_id": id}, value, options.Replace().SetUpsert(true),
	)
	return err
}

func (m *MongoDB) Read(table, id string, value any) (err error) {
	return m.Db.Collection(table).FindOne(
		context.Background(), bson.M{
			"_id": id,
		},
	).Decode(value)
}
