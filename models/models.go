package models

type Book struct {
	Name        string `json:"name" bson:"name"`
	Auther      string `json:"auther" bson:"auther"`
	IsAvailable bool   `json:"is_available" bson:"is_available"`
}
