/*
 * File Name main.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	network "github.com/aditya37/backend-jobs/api/Model/Network"

	repository "github.com/aditya37/backend-jobs/api/Repository/Employe"
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
	repo := repository.NewEmployeImpl(dbConn)
	x := repo.GetEmployeById(2)

	for _,hasil := range x {
		jsonValue, _ := json.MarshalIndent(hasil,""," ")
		fmt.Printf(string(jsonValue))
	}
	// jsonValue, _ := json.MarshalIndent(x,""," ")
	// fmt.Printf(string(jsonValue))
	
	// service := service.NewRegionService(repo)

	// fmt.Println(service.GetCountry())
}