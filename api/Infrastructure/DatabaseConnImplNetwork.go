/*
 * File Name DatabaseConnImplNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable search_path=public password=%s", host, port, username, dbname, password)
	s.Database,err = sqlx.Connect("postgres",URL)
	if err != nil {
		return nil,err
	}
	return s.Database,nil
}

func (s *_databaseConnection) DatabaseMigrate() error {
	DBDriver,err := postgres.WithInstance(s.Database.DB,&postgres.Config{})
	if err != nil {
		return err
	}
	migrate,err := migrate.NewWithDatabaseInstance("file://Database","postgres",DBDriver)
	if err != nil {
		return err
	}
	if err := migrate.Up(); err != nil {
		return err
	}
	return nil
}
