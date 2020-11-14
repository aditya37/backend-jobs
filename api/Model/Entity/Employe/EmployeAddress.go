/*
 * File Name EmployeAddress.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAddress struct {
	CountryName  string `db:"country_name" json:"country_name,omitempty"`
	ProvinceName string `db:"province_name" json:"province_name,omitempty"`
	DistrictName string `db:"district_name" json:"district_name,omitempty"`
	Address_1 	 string `db:"address_1" json:"address_1,omitempty"`
	Address_2 	 string `db:"address_2" json:"address_2,omitempty"`
	PostalCode   int	 `db:"postal_code" json:"postal_code,omitempty"`
	EmployeId	 int64  `db:"employe_id" json:"-"`
}