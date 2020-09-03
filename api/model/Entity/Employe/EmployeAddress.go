/*
 * File Name EmployeAddress.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

import (
	model "github.com/aditya37/backend-jobs/api/model/Entity"
)

type EmployeAddress struct {
	IdCountry	int `gorm:"type:SMALLINT;" json:"idCountry"`
	IdProvince	int `gorm:"type:SMALLINT;" json:"idProvince"`
	IdDistrict  int `gorm:"type:SMALLINT;" json:"idDistrict"`
	Address_1 	string `gorm:"type:VARCHAR(100);" json:"address_1"`
	Address_2 	string `gorm:"type:VARCHAR(100);" json:"address_2"`
	PostalCode  int	   `gorm:"type:SMALLINT" json:"postalCode"`
	EmployeId	int	   `gorm:"type:SMALLINT;" json:"IdEmploye"`
	Country 	model.Country  `gorm:"foreignkey:IdCountry"`
	Province 	model.Province `gorm:"foreignkey:IdProvince"`
	District 	model.District `gorm:"foreignkey:IdDistrict"`
	EmployeAccount EmployeAccount `gorm:"foreignkey:EmployeId"`
}