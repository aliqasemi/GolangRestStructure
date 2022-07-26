package main

import (
	"basical-app/config"
	"basical-app/router"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.Config{}
	err := config.SetConfig()
	if err != nil {
		fmt.Println(err)
	}
	echo := echo.New()
	router.SetRout(echo)

	echo.Logger.Fatal(echo.Start(":" + config.Server.Port))
}
