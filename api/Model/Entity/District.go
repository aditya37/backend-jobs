/*
 * File Name District.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package Model

type District struct {
	DistrictName string `gorm:gorm:"type:varchar(50);primary_key;" json:districtName`
}