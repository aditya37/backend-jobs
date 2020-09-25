/*
 * File Name EmployeAddress.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAddress struct {
	CountryName  string `gorm:"type:varchar(50);" json:"country_name"`
	ProvinceName string `gorm:"type:varchar(50);" json:"province_name"`
	DistrictName string `gorm:"type:varchar(50);" json:"district_name"`
	Address_1 	string `gorm:"type:VARCHAR(100);" json:"address_1"`
	Address_2 	string `gorm:"type:VARCHAR(100);" json:"address_2"`
	PostalCode  int	   `gorm:"type:INT;" json:"postal_code"`
	EmployeId	int64	`gorm:"type:INTEGER;"json:"-"`
}