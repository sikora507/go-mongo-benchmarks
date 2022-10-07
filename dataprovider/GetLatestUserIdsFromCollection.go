package dataprovider

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	UserId string `bson:"userId"`
}
type userIdContainer struct {
	UserId string `bson:"userId"`
	Users  []user `bson:"users"`
}

func GetLatestUserIdsFromCollection(collection *mongo.Collection, howMany int64) []string {
	ctx := context.TODO()
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(howMany)
	cursor, err := collection.Find(ctx, filter, opts)
	var users []userIdContainer
	// check for errors in the conversion
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}
	var result []string
	for _, v := range users {
		if v.UserId == "" {
			result = append(result, v.Users[0].UserId)
		} else {
			result = append(result, v.UserId)
		}
	}
	return result
}
