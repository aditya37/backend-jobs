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
	FirstName 	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Birth		time.Time `json:"birth"`
	BirthPlace 	string	`json:"birthPlace"`
	IsMale		string	`json:"isMale"`
	Phone		int		`json:"phone"`
	About		string	`json:"about"`
	EmployeId	int		`json:"IdEmploye"`
}