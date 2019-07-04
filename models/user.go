package models

type User struct {
	ID        ID     `json:"ID"`
	FirstName string `json:"first_name", bson:"first_name"`
	LastName  string `json:"last_name", bson:"last_name"`
}
