package dataprovider

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDataFromEmbeddedUsers(collection *mongo.Collection, userIds []string) []UserAccountDto {
	ctx := context.TODO()
	matchStage := bson.D{{"$match", bson.D{{"users.userId", bson.D{{"$in", userIds}}}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$users"}}}}
	matchStage2 := bson.D{{"$match", bson.D{{"users.userId", bson.D{{"$in", userIds}}}}}}
	projectStage := bson.D{{"$project", bson.D{
		{"accountId", 1},
		{"userId", "$users.userId"},
		{"tenantId", "$users.tenantId"},
		{"deleted", "$users.deleted"},
		{"email", 1},
		{"givenName", 1},
		{"familyName", 1},
	}}}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindStage, matchStage2, projectStage})
	// convert the cursor result to bson
	var results []UserAccountDto
	// check for errors in the conversion
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	return results
}
