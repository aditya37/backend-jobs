/*
 * File Name EmployeAccount.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package Model

import (
	"time"
)

type EmployeAccount struct {
	Id		 int	`gorm:"type:SMALLINT;primary_key" json:"idEmploye"`
	Username string `gorm:"type:VARCHAR(12)" json:"username"`
	Password string `gorm:"type:CHAR(16)" json:"password"`
	PhotoProfile string `gorm:"type:CHAR(100); json:"photoProfile"`
	RefreshToken string `gorm:"VARCHAR(20)" json:"refreshToken"`
	DateCreate time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"DateCreate"`
	DateUpdate time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"DateUpdate"`
	EmployeExperience  []EmployeExperience `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id"`
	EmployeAttachments EmployeAttachment  `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id"`
	EmployeAddress	   EmployeAddress	  `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id"`
	EmployeEducations  []EmployeEducation `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id"`
	EmployeDatas 	   EmployeData		  `gorm:"foreignKey:EmployeId;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;NOT NULL;references:id"`
}