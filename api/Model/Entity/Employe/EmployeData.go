/*
 * File Name EmployeData.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeData struct {
	FirstName 	string	`db:"first_name" json:"first_name,omitempty"`
	LastName	string	`db:"last_name" json:"last_name,omitempty"`
	Birth		string  `json:"birth,omitempty"`
	BirthPlace 	string	`db:"birth_place" json:"birth_place,omitempty"`
	IsMale		string	`db:"is_male" json:"isMale,omitempty"`
	Phone		int		`json:"phone,omitempty"`
	About		string	`json:"about,omitempty"`
	EmployeId 	int64 	`db:"employe_id" json:"-"`
}