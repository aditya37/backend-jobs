/*
 * File Name EmployeSocial.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeSocial struct {
	PortofolioLink string `json:"portofolioLink"`
	GithubLink 	   string `json:"githubLink"`
	LinkedinLink   string `json:"linkedinLink"`
	BlogLink 	   string `json:"blogLink"`
	TwitterLink    string `json:"twitterLink"`
	EmployeId	   int	  `json:"IdEmploye"`
}
