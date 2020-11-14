/*
 * File Name EmployeAccount.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

import (
	"time"
)

type EmployeAccount struct {
	Id		 int64	`db:"id" json:"idEmploye"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password,omitempty"`
	Email 	 string `db:"email" json:"email"`
	PhotoProfile string `db:"photo_profile" json:"photo_profile"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	IsActive 	 string `db:"is_active" json:"is_active"`
	DateCreate 	 time.Time `db:"date_create" json:"date_create"`
	DateUpdate 	 time.Time `db:"date_update" json:"date_update"`
	EmployeData 	  *EmployeData `json:"employe_datas"`
	EmployeEducation  *[]EmployeEducation `json:"employe_educations"`
	EmployeExperience *[]EmployeExperience `json:"employe_experiences"`
	EmployeAddress 	  *EmployeAddress `json:"employe_addresses"`
	EmployeAttachment *EmployeAttachment `json:"employe_attachments"`
	EmployeSocial 	  *EmployeSocial `json:"employe_socials"`
}
