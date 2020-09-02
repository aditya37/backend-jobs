/*
 * File Name EmployeAddress.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAddress struct {
	IdCountry	int `gorm:"type:SMALLINT;foreign_key;" json:"idCountry"`
	IdProvince	int `gorm:"type:SMALLINT;foreign_key;" json:"idProvince"`
	IdDistrict  int `gorm:"type:SMALLINT;foreign_key;" json:"idDistrict"`
	Address_1 	string `gorm:"type:VARCHAR(100);" json:"address_1"`
	Address_2 	string `gorm:"type:VARCHAR(100);" json:"address_2"`
	PostalCode  int	   `gorm:"type:SMALLINT" json:"postalCode"`
	EmployeId	int	   `gorm:"type:SMALLINT;foreign_key;" json:"IdEmploye"`
}