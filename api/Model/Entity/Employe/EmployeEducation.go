/*
 * File Name EmployeEducation.go
 * Created on Sun Sep 13 2020
 *
 * Copyright (c) 2020
 */

package model

import "time"

type EmployeEducation struct {
	InstitutionName	string `gorm:"type:VARCHAR(20);" json:"institution_name"`
	Degree			string `gorm:"type:VARCHAR(20);" json:"degree"`
	Certificate		string `gorm:"type:CHAR(20);" 	 json:"certificate_link"`
	IsActive 		string `gorm:"type:CHAR(5);" json:"IsActive"`
	StartEducation 	time.Time `gorm:"type:DATE;" json:"start_education"`
	EndEducation 	time.Time `gorm:"type:DATE;" json:"end_education"`
	EmployeId 		int64 `gorm:"type:INTEGER;NOT NULL;" json:"-"`
}