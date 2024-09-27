package main

import (
	"apartment_rent/configs"
	"apartment_rent/db"
	_ "apartment_rent/docs"
	"apartment_rent/logger"
	"apartment_rent/pkg/controllers"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
)

// @title Apartment_rent API
// @version 1.0
// @description API Server for Apartment_rent Application
// @host localhost:8181
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
	}

	err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}

	err = logger.Init()
	if err != nil {
		panic(err)
	}

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	err = controllers.RunRoutes()
	if err != nil {
		return
	}
}
