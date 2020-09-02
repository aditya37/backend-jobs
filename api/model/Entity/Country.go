/*
 * File Name Country.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type Country struct {
	// gorm.Model
	IdCountry 	int		`gorm:"type:smallint;primary_key" json:idCountry`
	CountryName string 	`gorm:"type:varchar(50)" json:countryName`
}