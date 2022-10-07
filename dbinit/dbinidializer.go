package dbinit

import (
	"context"
	"fmt"
	"go-mongo-benchmarks/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const databaseName = "UserLedgerBenchmark"
const accountsCollectionName = "accounts"
const usersCollectionName = "users"
const accountsWithUsersCollectionName = "accountsWithUsers"
const flattenedAccountsWithUsersCollectionName = "flattenedAccountsWithUsers"
const connectionString = "mongodb://admin:password@localhost:27018"

func InitDb(howManyAccounts int, howManyUsersPerAccount int) (*mongo.Collection, *mongo.Collection, *mongo.Collection, *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	client.Database(databaseName).Collection(accountsCollectionName).Drop(ctx)
	client.Database(databaseName).Collection(usersCollectionName).Drop(ctx)
	client.Database(databaseName).Collection(accountsWithUsersCollectionName).Drop(ctx)
	client.Database(databaseName).Collection(flattenedAccountsWithUsersCollectionName).Drop(ctx)

	client.Database(databaseName).CreateCollection(ctx, accountsCollectionName)
	client.Database(databaseName).CreateCollection(ctx, usersCollectionName)
	client.Database(databaseName).CreateCollection(ctx, accountsWithUsersCollectionName)
	client.Database(databaseName).CreateCollection(ctx, flattenedAccountsWithUsersCollectionName)

	// setup accounts and users
	accountsCollection := client.Database(databaseName).Collection(accountsCollectionName)
	usersCollection := client.Database(databaseName).Collection(usersCollectionName)
	for a := 0; a < howManyAccounts; a++ {
		accountId := utils.Guid()
		account := bson.D{
			{"accountId", accountId},
			{"email", fmt.Sprintf("%s%s", "email", utils.Guid())},
			{"givenName", fmt.Sprintf("%s%s", "givenName", utils.Guid())},
			{"familyName", fmt.Sprintf("%s%s", "familyName", utils.Guid())},
		}
		accountsCollection.InsertOne(ctx, account)
		var users []interface{}
		for u := 0; u < howManyUsersPerAccount; u++ {
			users = append(users, bson.D{
				{"accountId", accountId},
				{"userId", utils.Guid()},
				{"tenantId", utils.Guid()},
				{"deleted", false},
			})
		}
		usersCollection.InsertMany(ctx, users)
	}

	// setup accounts with internal users array
	accountsWithUsersCollection := client.Database(databaseName).Collection(accountsWithUsersCollectionName)
	var accountsWithUsers []interface{}
	for a := 0; a < howManyAccounts; a++ {
		var users []interface{}
		for u := 0; u < howManyUsersPerAccount; u++ {
			users = append(users, bson.D{
				{"userId", utils.Guid()},
				{"tenantId", utils.Guid()},
				{"deleted", false},
			})
		}
		account := bson.D{
			{"accountId", utils.Guid()},
			{"email", fmt.Sprintf("%s%s", "email", utils.Guid())},
			{"givenName", fmt.Sprintf("%s%s", "givenName", utils.Guid())},
			{"familyName", fmt.Sprintf("%s%s", "familyName", utils.Guid())},
			{"users", users},
		}
		accountsWithUsers = append(accountsWithUsers, account)
	}
	accountsWithUsersCollection.InsertMany(ctx, accountsWithUsers)

	// setup flattened data structure
	flattenedAccountsWithUsersCollection := client.Database(databaseName).Collection(flattenedAccountsWithUsersCollectionName)
	var flattenedAccountsWithUsers []interface{}
	for a := 0; a < howManyAccounts; a++ {
		accountId := utils.Guid()
		email := fmt.Sprintf("%s%s", "email", utils.Guid())
		givenName := fmt.Sprintf("%s%s", "givenName", utils.Guid())
		familyName := fmt.Sprintf("%s%s", "familyName", utils.Guid())
		for u := 0; u < howManyUsersPerAccount; u++ {
			account := bson.D{
				{"accountId", accountId},
				{"email", email},
				{"givenName", givenName},
				{"familyName", familyName},
				{"userId", utils.Guid()},
				{"tenantId", utils.Guid()},
				{"deleted", false},
			}
			flattenedAccountsWithUsers = append(flattenedAccountsWithUsers, account)
		}
	}
	flattenedAccountsWithUsersCollection.InsertMany(ctx, flattenedAccountsWithUsers)

	// setup indexes
	indexModelAccounts := mongo.IndexModel{
		Keys: bson.M{
			"accountId": 1,
		}, Options: nil,
	}
	accountsCollection.Indexes().CreateOne(ctx, indexModelAccounts)

	indexModelUsers := mongo.IndexModel{
		Keys: bson.M{
			"userId": 1,
		}, Options: nil,
	}
	usersCollection.Indexes().CreateOne(ctx, indexModelUsers)

	indexModelAccountsWithUsers := []mongo.IndexModel{
		{
			Keys: bson.M{
				"accountId": 1,
			}, Options: nil,
		},
		{
			Keys: bson.M{
				"users.userId": 1,
			}, Options: nil,
		},
	}
	accountsWithUsersCollection.Indexes().CreateMany(ctx, indexModelAccountsWithUsers)

	indexModelFlattenedAccountsWithUsers := []mongo.IndexModel{
		{
			Keys: bson.M{
				"accountId": 1,
			}, Options: nil,
		},
		{
			Keys: bson.M{
				"userId": 1,
			}, Options: nil,
		},
	}
	flattenedAccountsWithUsersCollection.Indexes().CreateMany(ctx, indexModelFlattenedAccountsWithUsers)

	return accountsCollection, usersCollection, accountsWithUsersCollection, flattenedAccountsWithUsersCollection
}
