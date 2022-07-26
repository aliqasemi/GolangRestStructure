package database

import (
	"basical-app/config"
	"fmt"
)

func init() {
	config := config.Config{}
	err := config.SetConfig()
	if err != nil {
		fmt.Println(err)
	}
	ConnectToMongo(config.Database.Mongo.Connection.Text)
}
