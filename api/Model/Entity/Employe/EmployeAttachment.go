/*
 * File Name EmployeAttachment.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAttachment struct {
	PortofolioFile string `gorm:"type:CHAR(100);" json:"portofolio_file"`
	ResumeFile 	   string `gorm:"type:CHAR(100);" json:"resume_file"`
	EmployeId	   int64  `gorm:"type:INTEGER;"json:"-"`
}