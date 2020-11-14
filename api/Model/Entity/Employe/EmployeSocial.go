/*
 * File Name EmployeSocial.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeSocial struct {
	PortofolioLink string `db:"portofolio_link" json:"portofolio_link,omitempty"`
	GithubLink 	   string `db:"github_link" json:"github_link,omitempty"`
	LinkedinLink   string `db:"linkedin_link" json:"linkedin_link,omitempty"`
	BlogLink 	   string `db:"blog_link" json:"blog_link,omitempty"`
	TwitterLink    string `db:"twitter_link" json:"twitter_link,omitempty"`
	EmployeId 	   int64 `db:"employe_id" json:"-"`
}
