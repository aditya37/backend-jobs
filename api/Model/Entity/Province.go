/*
 * File Name Province.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package Model

type Province struct {
	ProvinceName string `gorm:"type:varchar(50);primary_key;" json:provinceName`
}