/*
 * File Name IDatabaseConnNetwork.go
 * Created on Fri Aug 21 2020
 *
 * Copyright (c) 2020
 */

package network

import (
	"gorm.io/gorm"
)

type DatabaseConnection interface {
	DatabaseConn(host,port,username,password,dbname string) (*gorm.DB,error)
	DatabaseMigrate()
}

 
 