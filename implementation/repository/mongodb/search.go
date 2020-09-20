package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/aeramu/nim-finder-server/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Search(keyword string) []entity.User {
	opt := options.Find()
	opt.SetSort(bson.D{{"status", -1}})

	cursor, err := r.nim.Find(context.TODO(), bson.D{{"$or", bson.A{
		bson.D{{"nama", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_jurusan", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_tpb", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
	}}}, opt)
	if err != nil {
		panic(err)
	}

	var users []entity.User
	cursor.All(context.TODO(), &users)

	return users
}
