/*
 * File Name EmployeAttachment.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package Model

type EmployeAttachment struct {
	PortofolioFile string `gorm:"type:CHAR(100);" json:"portofolioFile"`
	ResumeFile 	   string `gorm:"type:CHAR(100);" json:"resumeFile"`
	EmployeId 	   int `json:"employeID"`
}