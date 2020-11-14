/*
 * File Name DatabaseConnImplNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type _databaseConnection struct {
	Database *sqlx.DB
}

func NewDatabaseConnection() DatabaseConnection {
	return &_databaseConnection{}
	
}

func (s *_databaseConnection) DatabaseConn(host,port,username,password,dbname string) (*sqlx.DB,error){
	var err error

	URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)
	s.Database,err = sqlx.Connect("postgres",URL)
	if err != nil {
		return nil,err
	}
	
	return s.Database,nil
}

// func (s *_databaseConnection) DatabaseMigrate() {
// 	// s.Database.Migrator().DropTable(&model.EmployeAccount{},&model.EmployeAddress{},&model.EmployeAttachment{},&model.EmployeData{},&model.EmployeExperience{},&model.EmployeSocial{},&model.EmployeEducation{},&region.Country{},&region.District{},&region.Province{})
// 	s.Database.AutoMigrate(&model.EmployeAccount{},&model.EmployeAddress{},&model.EmployeAttachment{},&model.EmployeData{},&model.EmployeExperience{},&model.EmployeSocial{},&model.EmployeEducation{},&region.Country{},&region.District{},&region.Province{})
// }
