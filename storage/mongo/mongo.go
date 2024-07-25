package mongodb

import (
	"context"
	"fmt"
	"log"

	u "github.com/dilshodforever/4-oyimtixon-game-service/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	Db    *mongo.Database
	Games u.GameStorage
}

func NewMongoConnecti0n() (u.InitRoot, error) {
	uri := fmt.Sprintf("mongodb://%s:%d",
    "mongo",
    27017,
  )
	clientOptions := options.Client().ApplyURI(uri).
	SetAuth(options.Credential{Username: "dilshod", Password: "root"})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error: Couldn't connect to the database.", err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("smm")

	return &MongoStorage{Db: db}, err
}

func (s *MongoStorage) Game() u.GameStorage {
	if s.Games == nil {
		s.Games = &GameStorage{s.Db}
	}
	return s.Games
}
