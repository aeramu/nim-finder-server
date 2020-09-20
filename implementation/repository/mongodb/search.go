package mongodb

import (
	"context"

	"github.com/aeramu/nim-finder-server/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Search(keyword string) []entity.User {
	cursor, err := r.nim.Find(context.TODO(), bson.D{{"$or", bson.A{
		bson.D{{"nama", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_jurusan", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_tpb", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
	}}})
	if err != nil {
		panic(err)
	}

	var users []entity.User
	cursor.All(context.TODO(), &users)

	return users
}
