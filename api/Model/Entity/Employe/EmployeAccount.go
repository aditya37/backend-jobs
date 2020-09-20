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
	Id		 int	`gorm:"type:SMALLINT;primary_key" json:"idEmploye"`
	Username string `gorm:"type:VARCHAR(12);" json:"username"`
	Password string `gorm:"type:CHAR(60);" json:"-"`
	Email 	 string `gorm:"type:VARCHAR(60);" json:"email"`
	PhotoProfile string `gorm:"type:CHAR(100);" json:"photo_profile"`
	RefreshToken string `gorm:"VARCHAR(20);" json:"refresh_token"`
	IsActive 	 string `gorm:"CHAR(5);" json:"isActive"`
	DateCreate time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"date_create"`
	DateUpdate time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"date_update"`
	EmployeExperience  []EmployeExperience `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id;" json:"employe_experiences"`
	EmployeAttachments EmployeAttachment  `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id;"  json:"employe_attachments"`
	EmployeAddress	   EmployeAddress	  `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id;"  json:"employe_address"`
	EmployeEducations  []EmployeEducation `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id;"  json:"employe_educations"`
	EmployeDatas 	   EmployeData		  `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id;"  json:"employe_datas"`
}