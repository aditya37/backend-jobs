/*
 * File Name EmployeData.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeData struct {
	FirstName 	string	`gorm:"type:VARCHAR(20);" json:"first_name"`
	LastName	string	`gorm:"type:VARCHAR(200);" json:"last_name"`
	Birth		string `gorm:"type:DATE;" json:"birth"`
	BirthPlace 	string	`gorm:"type:CHAR(20);" json:"birth_place"`
	IsMale		string	`gorm:"type:CHAR(5);" json:"isMale"`
	Phone		int		`gorm:"type:INT;" json:"phone"`
	About		string	`gorm:"type:VARCHAR(200);" json:"about"`
	EmployeId 	int64 	`gorm:"type:INTEGER;NOT NULL;" json:"-"`
}