/*
 * File Name EmployeWork.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

import "time"

type EmployeExperience struct {
	CompanyName string `json:"companyName"`
	JobTitle 	string `json:"jobTitle"`
	JobDesc 	string `json:"jobDesc"`
	IsActive 	string `json:"isActive"`
	StartWork 	time.Time `json:"startWork"`
	EndWork 	time.Time `json:"endWork"`
	EmployeId	int		`json:"IdEmploye"`
}