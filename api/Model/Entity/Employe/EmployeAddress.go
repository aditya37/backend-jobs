/*
 * File Name EmployeAddress.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package Model

type EmployeAddress struct {
	CountryName  string `gorm:"type:varchar(50);" json:countryName`
	ProvinceName string `gorm:"type:varchar(50);" json:provinceName`
	DistrictName string `gorm:"type:varchar(50);" json:districtName`
	Address_1 	string `gorm:"type:VARCHAR(100);" json:"address_1"`
	Address_2 	string `gorm:"type:VARCHAR(100);" json:"address_2"`
	PostalCode  int	   `gorm:"type:SMALLINT" json:"postalCode"`
	EmployeId	int	   `gorm:"type:SMALLINT;"json:"IdEmploye"`
}