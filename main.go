/*
 * File Name main.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aditya37/backend-jobs/api/network"
	"github.com/aditya37/backend-jobs/api/repository"
	"github.com/joho/godotenv"
)

var (
	connection network.DatabaseConnection = network.NewDatabaseConnection()
)

func main() {
	env := godotenv.Load()
	if env != nil {
		log.Panic("Error .env file not found")
	}
	dbConn,err := connection.DatabaseConn(os.Getenv("DBHOST"),os.Getenv("DBPORT"),os.Getenv("DBUSER"),os.Getenv("DBPASSWORD"),os.Getenv("DBNAME"))
	if err != nil {
		fmt.Println("Database connection error ",err)
	}
	connection.DatabaseMigrate()
	repo := repository.NewRegionImpl(dbConn)
	repo.GetCountry()
}