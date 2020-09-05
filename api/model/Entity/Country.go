/*
 * File Name Country.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type Country struct {
	CountryName string 	`gorm:"type:varchar(50);primary_key;" json:countryName`
}