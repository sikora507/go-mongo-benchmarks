package dataprovider

import "testing"

func BenchmarkGetDataFromFlattenedStructure(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetDataFromFlattenedStructure(flattenedAccountsWithUsersCollection, userIdsFlattened)
	}
}
