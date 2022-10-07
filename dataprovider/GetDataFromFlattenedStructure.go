package dataprovider

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDataFromFlattenedStructure(collection *mongo.Collection, userIds []string) []UserAccountDto {
	ctx := context.TODO()
	filter := bson.D{{"userId", bson.D{{"$in", userIds}}}}
	cursor, err := collection.Find(ctx, filter)
	// convert the cursor result to bson
	var results []UserAccountDto
	// check for errors in the conversion
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results
}
