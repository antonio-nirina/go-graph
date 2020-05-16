package model

type User struct {
	ID        string `json:"_id"`
	Addresse  string `json:"addresse"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
