/*
 * File Name EmployeEducation.go
 * Created on Sun Sep 13 2020
 *
 * Copyright (c) 2020
 */

package model
type EmployeEducation struct {
	InstitutionName	string `db:"institution_name" json:"institution_name,omitempty"`
	Degree			string `db:"degree" json:"degree,omitempty"`
	IsActive 		string `db:"is_active" json:"IsActive,omitempty"`
	StartEducation 	string `db:"start_education" json:"start_education,omitempty"`
	EndEducation 	string `db:"end_education" json:"end_education,omitempty"`
	EmployeId 		int64  `db:"employe_id" json:"-"`
}