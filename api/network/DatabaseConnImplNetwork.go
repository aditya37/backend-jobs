/*
 * File Name DatabaseConnImplNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package network

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

type _databaseConnection struct {
	Database *sql.DB
}

func NewDatabaseConnection() DatabaseConnection {
	return &_databaseConnection{}
}

func (s *_databaseConnection) DatabaseConn(host,port,username,password,dbname string) (*sql.DB,error){
	
	var err error
	URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)
	s.Database,err = sql.Open("postgres",URL)
	if err != nil {
		return nil,err
	}

	if err := s.Database.Ping(); err != nil {
		return nil,err
	}

	return s.Database,err
}
