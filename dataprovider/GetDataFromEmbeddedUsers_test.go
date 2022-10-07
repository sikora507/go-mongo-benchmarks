package dataprovider

import "testing"

func BenchmarkGetDataFromEmbeddedUsers(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetDataFromEmbeddedUsers(accountsWithUsersCollection, userIdsForEmbedded)
	}
}
