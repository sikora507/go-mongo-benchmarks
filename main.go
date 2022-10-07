package main

import (
	"fmt"
	"go-mongo-benchmarks/dataprovider"
	"go-mongo-benchmarks/dbinit"
)

const howManyAccounts = 10000
const howManyUsersPerAccount = 50

func main() {
	_, usersCollection, accountsWithUsersCollection, flattenedAccountsWithUsersCollection := dbinit.InitDb(howManyAccounts, howManyUsersPerAccount)

	userIdsFlattened := dataprovider.GetLatestUserIdsFromCollection(flattenedAccountsWithUsersCollection, 40)
	userIdsForJoin := dataprovider.GetLatestUserIdsFromCollection(usersCollection, 40)
	userIdsForEmbedded := dataprovider.GetLatestUserIdsFromCollection(accountsWithUsersCollection, 40)

	dataFromFlattened := dataprovider.GetDataFromFlattenedStructure(flattenedAccountsWithUsersCollection, userIdsFlattened)
	dataFromJoining := dataprovider.GetDataByJoiningDocuments(usersCollection, userIdsForJoin)
	dataFromEmbedded := dataprovider.GetDataFromEmbeddedUsers(accountsWithUsersCollection, userIdsForEmbedded)

	fmt.Println("Data from flat structure:")
	fmt.Println(dataFromFlattened[0:10])
	fmt.Println("Data from joined structure:")
	fmt.Println(dataFromJoining[0:10])
	fmt.Println("Data from embedded structure:")
	fmt.Println(dataFromEmbedded[0:10])
}
