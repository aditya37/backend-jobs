/*
 * File Name EmployeWork.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

import (
	"time"
)

type EmployeExperience struct {
	CompanyName string `gorm:"type:VARCHAR(20);" json:"companyName"`
	JobTitle 	string `gorm:"type:VARCHAR(20);" json:"jobTitle"`
	JobDesc 	string `gorm:"type:VARCHAR(200);" json:"jobDesc"`
	IsActive 	string `gorm:"type:CHAR(5);" json:"isActive"`
	StartWork 	time.Time `gorm:"type:DATE" json:"startWork"`
	EndWork 	time.Time `gorm:"type:DATE;" json:"endWork"`
	EmployeId 	int64 	  `gorm:"type:INTEGER;NOT NULL;" json:"-"`
}