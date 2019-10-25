package util

// Auth : a string that represent authorization
type Auth string

//Role : A relation ship about authorization
type Role uint32

const (
	// Nobody : Anyone, even have not logged in; the lowest auth, basically no auth;
	Nobody = 0
	//Acquaintance : People who know each other(follow each other); can have access to not sensitive data;
	Acquaintance = 10
	// Friend : People who have close relation ship(have added each other); can have access to detailed data;
	Friend = 20
	// Manager : People who manage that user; can have RW rights to public data;
	Manager = 30
)

//UserID : the uniqe uint64 integer to identify user
type UserID uint64

