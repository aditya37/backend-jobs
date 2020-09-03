/*
 * File Name EmployeData.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

import (
	"time"
)

type EmployeData struct {
	FirstName 	string	`gorm:"type:VARCHAR(20);" json:"firstName"`
	LastName	string	`gorm:"type:VARCHAR(200);" json:"lastName"`
	Birth		time.Time `gorm:"type:DATE;" json:"birth"`
	BirthPlace 	string	`gorm:"type:CHAR(20);" json:"birthPlace"`
	IsMale		string	`gorm:"type:CHAR(5);" json:"isMale"`
	Phone		int		`gorm:"type:INT;" json:"phone"`
	About		string	`gorm:"type:VARCHAR(200);" json:"about"`
	EmployeId	int		`gorm:"type:SMALLINT;" json:"IdEmploye"`
	EmployeAccount EmployeAccount `gorm:"foreignkey:EmployeId"`
}