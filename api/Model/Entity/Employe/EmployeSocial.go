/*
 * File Name EmployeSocial.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeSocial struct {
	PortofolioLink string `gorm:"type:CHAR(100)" json:"portofolioLink"`
	GithubLink 	   string `gorm:"type:CHAR(100)" json:"githubLink"`
	LinkedinLink   string `gorm:"type:CHAR(100)" json:"linkedinLink"`
	BlogLink 	   string `gorm:"type:CHAR(100)" json:"blogLink"`
	TwitterLink    string `gorm:"type:CHAR(100)" json:"twitterLink"`
	EmployeId 	   int64 `gorm:"type:INTEGER;NOT NULL;" json:"-"`
}
