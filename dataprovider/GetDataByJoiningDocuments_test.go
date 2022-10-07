package dataprovider

import "testing"

func BenchmarkGetDataByJoiningDocuments(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetDataByJoiningDocuments(usersCollection, userIdsForJoin)
	}
}
