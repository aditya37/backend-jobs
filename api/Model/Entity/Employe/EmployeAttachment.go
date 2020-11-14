/*
 * File Name EmployeAttachment.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAttachment struct {
	PortofolioFile string `db:"portofolio_file" json:"portofolio_file"`
	ResumeFile 	   string `db:"resume_file" json:"resume_file"`
	ResumeObject   string `db:"resume_object"  json:"resume_object,omitempty"`
	PortofolioObject string `db:"portofolio_object" json:"portofolio_object,omitempty"`
	EmployeId	   int64  `db:"employe_id" json:"-"`
}