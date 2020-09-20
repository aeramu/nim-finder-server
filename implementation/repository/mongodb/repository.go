package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	client *mongo.Client
	nim    *mongo.Collection
}
