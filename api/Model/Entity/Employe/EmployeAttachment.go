/*
 * File Name EmployeAttachment.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAttachment struct {
	PortofolioFile string `gorm:"type:VARCHAR(255);" json:"portofolio_file"`
	ResumeFile 	   string `gorm:"type:VARCHAR(255);" json:"resume_file"`
	ResumeObject   string `gorm:"type:VARCHAR(50);"  json:"resume_object"`
	PortofolioObject string `gorm:"type:VARCHAR(50);"  json:"portofolio_object"`
	EmployeId	   int64  `gorm:"type:INTEGER;"json:"-"`
}