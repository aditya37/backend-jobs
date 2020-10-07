/*
 * File Name main.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package main

import (
	"log"
	"os"

	controller "github.com/aditya37/backend-jobs/api/Controller/Employe"
	infrastructure "github.com/aditya37/backend-jobs/api/Infrastructure"
	repository "github.com/aditya37/backend-jobs/api/Repository/Employe"
	service "github.com/aditya37/backend-jobs/api/Service/Employe"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	router infrastructure.IEchoRouter = infrastructure.NewEchoRouter(echo.New())
	database infrastructure.DatabaseConnection = infrastructure.NewDatabaseConnection()
	redisCache infrastructure.IRedisConn = infrastructure.NewRedisConn("localhost:6379","",0,900)
	firebase infrastructure.IFireStorage = infrastructure.NewFireStorageImpl()
)

func main() {
	LoadEnv := godotenv.Load()
	if LoadEnv != nil {
		log.Fatalln(LoadEnv)
	}
	
	ConnectDB,err := database.DatabaseConn(os.Getenv("DBHOST"),os.Getenv("DBPORT"),os.Getenv("DBUSER"),os.Getenv("DBPASSWORD"),os.Getenv("DBNAME"))
	if err != nil {
		log.Panic(err)
	}
	database.DatabaseMigrate()

	EmployeRepo 	:= repository.NewEmployeImpl(ConnectDB)
	EmployeService  := service.NewEmployeService(EmployeRepo)
	EmployeController := controller.NewEmployeController(EmployeService,redisCache,firebase)
	
	// Verify route
	router.Get("/verify/employe/verifyEmail",EmployeController.VerifyEmail)
	router.Post("/verify/employe/refreshEmailVerify",EmployeController.RefreshEmailVerify)
	
	// Subroute or group Route
	EmployeRoutes := router.RouteGroup("employes")
	EmployeRoutes.POST("/signup",EmployeController.RegisterEmploye)
	EmployeRoutes.POST("/signin",EmployeController.LoginEmploye)
	EmployeRoutes.GET("/:id",EmployeController.GetEmployeById)
	
	router.ErrorHandler() // Middleware error handling
	
	router.RouterLogger() // Router Logging
	router.StartServer(":3000") // Start server
}