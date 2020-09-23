package mongodb

import (
	"context"

	"github.com/aeramu/nim-finder-server/usecase/adapter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//New repository
func New() adapter.Repository {
	client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb+srv://admin:admin@qiup-wrbox.mongodb.net/",
	))
	return &repository{
		client: client,
		nim:    client.Database("itb").Collection("nim"),
	}
}

type repository struct {
	client *mongo.Client
	nim    *mongo.Collection
}
