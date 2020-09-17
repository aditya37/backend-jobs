/*
 * File Name EmployeEducation.go
 * Created on Sun Sep 13 2020
 *
 * Copyright (c) 2020
 */

package Model

import "time"

type EmployeEducation struct {
	InstitutionName	string `gorm:"type:VARCHAR(20);" json:"institutionName"`
	Degree			string `gorm:"type:VARCHAR(20);" json:"degree"`
	Certificate		string `gorm:"type:CHAR(20);" 	 json:"certificateLink"`
	IsActive 		string `gorm:"type:CHAR(5);" json:"IsActive"`
	StartEducation 	time.Time `gorm:"type:DATE;" json:"startEducation"`
	EndEducation 	time.Time `gorm:"type:DATE;" json:"endEducation"`
	EmployeId 		int `gorm:"type:SMALLINT;NOT NULL;" json:"employeID"`
}