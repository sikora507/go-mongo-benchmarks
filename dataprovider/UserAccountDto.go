package dataprovider

type UserAccountDto struct {
	AccountId  string `bson:"accountId"`
	UserId     string `bson:"userId"`
	TenantId   string `bson:"tenantId"`
	Deleted    bool   `bson:"deleted"`
	Email      string `bson:"email"`
	GivenName  string `bson:"givenName"`
	FamilyName string `bson:"familyName"`
}
