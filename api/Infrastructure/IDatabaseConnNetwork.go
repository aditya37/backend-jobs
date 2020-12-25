/*
 * File Name IDatabaseConnNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"github.com/jmoiron/sqlx"
)

type DatabaseConnection interface {
	DatabaseConn(host,port,username,password,dbname string) (*sqlx.DB,error)
	DatabaseMigrate() error
} 