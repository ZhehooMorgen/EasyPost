package util

// Auth : a string that represent authorization
type Auth string

//Role : A relation ship about authorization
type Role uint32

const (
	// Nobody : Anyone, even have not logged in; the lowest auth, basicly no auth;
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

//UserInfo : the data structure to describe user
type UserInfo struct {
	//ID : the uniqe uint64 integer to identify user
	ID UserID `json:"id"`
	//Name : the users display name, can be changed frequently
	Name string `json:"userName"`
	//password : should never get accessable
	password string
	//Phone ：phone number or other contact method of a user, can not leak
	Phone string `json:"phone"`
}
