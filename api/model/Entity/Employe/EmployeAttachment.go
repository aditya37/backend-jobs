/*
 * File Name EmployeAttachment.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAttachment struct {
	PortofolioFile string `json:"portofolioFile"`
	ResumeFile 	   string `json:"resumeFile"`
	EmployeId	   int	  `json:"IdEmploye"`
}