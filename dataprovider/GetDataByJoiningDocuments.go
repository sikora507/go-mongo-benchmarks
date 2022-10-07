package dataprovider

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDataByJoiningDocuments(usersCollection *mongo.Collection, userIds []string) []UserAccountDto {
	ctx := context.TODO()
	matchStage := bson.D{{"$match", bson.D{{"userId", bson.D{{"$in", userIds}}}}}}
	lookupStage := bson.D{{"$lookup", bson.D{
		{"from", "accounts"},
		{"localField", "accountId"},
		{"foreignField", "accountId"},
		{"as", "account"},
	}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$account"}}}}
	projectStage := bson.D{{"$project", bson.D{
		{"accountId", 1},
		{"userId", 1},
		{"tenantId", 1},
		{"deleted", 1},
		{"email", "$account.email"},
		{"givenName", "$account.givenName"},
		{"familyName", "$account.familyName"},
	}}}
	cursor, err := usersCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, projectStage})
	// convert the cursor result to bson
	var results []UserAccountDto
	// check for errors in the conversion
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	return results
}
