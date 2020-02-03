package entity

//User defines a single user
type User struct {
	ID       int
	FullName string
	UserName string
	Password string
	Picture  string
	Address  string
}

//Place defines a single Place
type Place struct {
	ID          int
	UserID      int
	Name        string
	Location    string
	Description string
}

//Picture defines a single picture
type Picture struct {
	ID      int
	PlaceID int
	Name    string
}
