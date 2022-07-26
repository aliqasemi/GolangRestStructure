package database

import (
	"context"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type MongoDb struct {
	client *mongo.Client
}

type singleton struct {
}

var (
	instance  *singleton
	clientVar MongoDb
)

func ConnectToMongo(uri string) MongoDb {
	if uri == "" {
		log.Fatal("You must set your 'database.mongo.connection.text' environmental variable in config.yaml.")
	}
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &singleton{}
		clientVar.client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	}
	return clientVar
}

func GetClient() MongoDb {
	return clientVar
}

func (db MongoDb) GetModelCollection(model string) *mongo.Collection {
	collection := db.client.Database("go").Collection(model)
	return collection
}
