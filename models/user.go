package models

type User struct {
	ID        string `json:"id" form:"id" query:"id" param:"id" bson:"_id,omitempty"`
	FirstName string `json:"first-name" form:"first-name" query:"first-name" param:"first-name" bson:"first-name,omitempty"`
	LastName  string `json:"last-name" form:"last-name" query:"last-name" param:"last-name" bson:"last-name,omitempty"`
	Age       int    `json:"age" form:"age" query:"age" param:"age" bson:"age,omitempty"`
}
