/*
 * File Name DatabaseConnImplNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package network

import (
	"fmt"

	model "github.com/aditya37/backend-jobs/api/model/Entity/Employe"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type _databaseConnection struct {
	Database *gorm.DB
}

func NewDatabaseConnection() DatabaseConnection {
	return &_databaseConnection{}
}

func (s *_databaseConnection) DatabaseConn(host,port,username,password,dbname string) (*gorm.DB,error){
	var err error
	URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)
	s.Database,err = gorm.Open(postgres.Open(URL),&gorm.Config{})
	if err != nil {
		return nil,err
	}	
	return s.Database,err
}

func (s *_databaseConnection) DatabaseMigrate() {
	s.Database.AutoMigrate(&model.EmployeAccount{},&model.EmployeAddress{},&model.EmployeAttachment{},&model.EmployeData{},&model.EmployeExperience{},&model.EmployeSocial{})
}