/*
 * File Name DatabaseConnImplNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"database/sql"
	"fmt"

	region "github.com/aditya37/backend-jobs/api/Model/Entity"
	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	_ "github.com/lib/pq"

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
	pqd,err:= sql.Open("postgres",URL)
	if err != nil {
		return nil,err
	}
	s.Database,err = gorm.Open(postgres.New(postgres.Config{Conn:pqd}),&gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil,err
	}	
	return s.Database,err
}

func (s *_databaseConnection) DatabaseMigrate() {
	// s.Database.Migrator().DropTable(&model.EmployeAccount{},&model.EmployeAddress{},&model.EmployeAttachment{},&model.EmployeData{},&model.EmployeExperience{},&model.EmployeSocial{},&model.EmployeEducation{},&region.Country{},&region.District{},&region.Province{})
	s.Database.AutoMigrate(&model.EmployeAccount{},&model.EmployeAddress{},&model.EmployeAttachment{},&model.EmployeData{},&model.EmployeExperience{},&model.EmployeSocial{},&model.EmployeEducation{},&region.Country{},&region.District{},&region.Province{})
}