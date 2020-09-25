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
	"time"

	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	"github.com/segmentio/ksuid"

	infra "github.com/aditya37/backend-jobs/api/Infrastructure"

	repository "github.com/aditya37/backend-jobs/api/Repository/Employe"
	service "github.com/aditya37/backend-jobs/api/Service/Employe"

	"github.com/joho/godotenv"
)

var (
	connection infra.DatabaseConnection = infra.NewDatabaseConnection()
	redisConn infra.IRedisConn = infra.NewRedisConn("localhost:6379","",0,10)
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
	EmployeRepo := repository.NewEmployeImpl(dbConn)
	Empl := service.NewEmployeService(EmployeRepo)
	
	account := &model.EmployeAccount{
		Username: "arifnyik",
		Password: "lymousin",
		Email: "aditya.krohman@gmail.com",
		PhotoProfile: "google.com",
		DateCreate: time.Now(),
		DateUpdate: time.Now(),
	}
	
	DoRegister,err := Empl.RegisterEmploye(account)
	if err != nil {
		log.Println(err.Error())
	}
	id := ksuid.New()
	err = redisConn.AddEmailVerify(DoRegister.Email,id.String())
	if err != nil {
		log.Println(err)
	}
	Empl.SendEmailVerify(DoRegister.Email,DoRegister.Username,id.String())
	redisKey,err := redisConn.VerifyEmail(DoRegister.Email)
	if err != nil {
		log.Println(err)
	}
	if len(redisKey) == 0 {
		fmt.Println("Cache Not Found",redisKey)
	}else{
		if redisKey == id.String() {
			EmployeRepo.EmployeEmailVerify(DoRegister.Email)
		}
	}

	// DoLogin,err := Empl.EmployeLogin("aditya","lymousin")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(DoLogin)
}