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

	model "github.com/aditya37/backend-jobs/api/model"
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
	repo := repository.NewRegionImpl(dbConn)
	negara,_ := repo.GetCountry()
	res := &model.SuccessResponse{
		Status: 1,
		Message: "Success",
		Result: negara,
	}
	b,_ := json.Marshal(res)
	fmt.Println(string(b))

}