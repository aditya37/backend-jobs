/*
 * File Name Province.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type Province struct {
	IdProvince 	 int 	`gorm:"type:smallint;primary_key" json:idProvince`
	ProvinceName string `gorm:"type:varchar(50)" json:provinceName`
}