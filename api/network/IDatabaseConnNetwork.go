/*
 * File Name IDatabaseConnNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package network

import (
	"database/sql"
)

type DatabaseConnection interface {
	DatabaseConn(host,port,username,password,dbname string) (*sql.DB,error)
}

 
 