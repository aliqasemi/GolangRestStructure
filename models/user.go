package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          string    `json:"id" form:"id" query:"id" param:"id" bson:"_id,omitempty"`
	FirstName   string    `json:"first-name" form:"first-name" query:"first-name" param:"first-name" bson:"first-name,omitempty"`
	LastName    string    `json:"last-name" form:"last-name" query:"last-name" param:"last-name" bson:"last-name,omitempty"`
	UserName    string    `json:"user-name" form:"user-name" query:"user-name" param:"user-name" bson:"user-name,omitempty"`
	Age         int       `json:"age" form:"age" query:"age" param:"age" bson:"age,omitempty"`
	Password    string    `json:"password" form:"password" query:"password" param:"password" bson:"password,omitempty"`
	PhoneNumber string    `json:"phone-number" form:"phone-number" query:"phone-number" param:"phone-number" bson:"phone-number,omitempty"`
	Email       string    `json:"email" form:"email" query:"email" param:"email" bson:"email,omitempty"`
	CreatedAt   time.Time `json:"created-at" form:"created-at" query:"created-at" param:"created-at" bson:"created-at,omitempty"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
