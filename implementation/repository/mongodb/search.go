package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/aeramu/nim-finder-server/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Search(keyword string, limit int) []entity.User {
	opt := options.Find().SetSort(bson.D{{"status", -1}}).SetLimit(int64(limit))

	search := bson.D{{"$or", bson.A{
		bson.D{{"nama", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_jurusan", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_tpb", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
	}}}

	cursor, err := r.nim.Find(context.TODO(), search, opt)
	if err != nil {
		panic(err)
	}

	var users []entity.User
	cursor.All(context.TODO(), &users)

	return users
}
