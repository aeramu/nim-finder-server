package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aeramu/nim-finder-server/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Search(keyword string, limit int, after string) entity.UserConnection {
	search := bson.D{{"$or", bson.A{
		bson.D{{"nama", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_jurusan", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
		bson.D{{"nim_tpb", bson.D{{"$regex", ".*" + keyword + ".*"}, {"$options", "i"}}}},
	}}}

	afterID, _ := primitive.ObjectIDFromHex(after)
	paginationStage := bson.D{{"$match", bson.D{{"_id", bson.D{{"$lt", afterID}}}}}}
	limitStage := bson.D{{"$limit", int64(limit + 1)}}
	sortStage := bson.D{{"$sort", bson.D{{"_id", -1}}}}
	matchStage := bson.D{{"$match", search}}

	cursor, err := r.nim.Aggregate(context.TODO(), mongo.Pipeline{matchStage, sortStage, paginationStage, limitStage})
	if err != nil {
		panic(err)
	}

	var models userModels
	cursor.All(context.TODO(), &models)

	var users entity.UserConnection
	if len(models) > limit {
		users = userConnection{models[:limit].entity(), true}
	} else {
		users = userConnection{models.entity(), false}
	}

	return users
}
