/*
 * File Name District.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type District struct {
	IdDistrict   int 	`gorm:"type:smallint;primary_key" json:idDistrict`
	DistrictName string `gorm:"type:varchar(50)" json:districtName`
}