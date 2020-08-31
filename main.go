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

	model "github.com/aditya37/backend-jobs/api/model"
	"github.com/aditya37/backend-jobs/api/network"
	"github.com/aditya37/backend-jobs/api/repository"
)

var (
	connection network.DatabaseConnection = network.NewDatabaseConnection()
)

func main() {
	dbConn,err := connection.DatabaseConn("localhost","5432","admin","lymousin","db_jobs")
	if err != nil {
		fmt.Println(err)
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