package dataprovider

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "UserLedgerBenchmark"
const usersCollectionName = "users"
const accountsWithUsersCollectionName = "accountsWithUsers"
const flattenedAccountsWithUsersCollectionName = "flattenedAccountsWithUsers"
const connectionString = "mongodb://admin:password@localhost:27018"

var client *mongo.Client
var userIdsFlattened []string
var userIdsForJoin []string
var userIdsForEmbedded []string
var usersCollection *mongo.Collection
var accountsWithUsersCollection *mongo.Collection
var flattenedAccountsWithUsersCollection *mongo.Collection

func Init() {
	if client == nil {

		client, _ = mongo.NewClient(options.Client().ApplyURI(connectionString))
		ctx := context.TODO()
		client.Connect(ctx)

		usersCollection = client.Database(databaseName).Collection(usersCollectionName)
		accountsWithUsersCollection = client.Database(databaseName).Collection(accountsWithUsersCollectionName)
		flattenedAccountsWithUsersCollection = client.Database(databaseName).Collection(flattenedAccountsWithUsersCollectionName)

		userIdsFlattened = GetLatestUserIdsFromCollection(flattenedAccountsWithUsersCollection, 40)
		userIdsForJoin = GetLatestUserIdsFromCollection(usersCollection, 40)
		userIdsForEmbedded = GetLatestUserIdsFromCollection(accountsWithUsersCollection, 40)
	}
}
