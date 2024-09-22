package serializers


type LoginUserIn struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}
type LoginUserOut struct {
	ID			uint64 `json:"id"`
	Email 		string `json:"email"`
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
}