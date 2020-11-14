/*
 * File Name EmployeWork.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeExperience struct {
	CompanyName string `db:"company_name" json:"company_name,omitempty"`
	JobTitle 	string `db:"job_title" json:"job_title,omitempty"`
	JobDesc 	string `db:"job_desc" json:"job_desc,omitempty"`
	IsActive 	string `db:"is_active" json:"is_active,omitempty"`
	StartWork 	string `db:"start_work" json:"start_work,omitempty"`
	EndWork 	string `db:"end_work" json:"end_work,omitempty"`
	EmployeId 	int64  `db:"employe_id" json:"-"`
}