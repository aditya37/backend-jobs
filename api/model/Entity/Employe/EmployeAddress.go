/*
 * File Name EmployeAddress.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

type EmployeAddress struct {
	IdCountry	int `json:"idCountry"`
	IdProvince	int `json:"idProvince"`
	IdDistrict  int `json:"idDistrict"`
	Address_1 	string `json:"address_1"`
	Address_2 	string `json:"address_2"`
	PostalCode  int	   `json:"postalCode"`
	EmployeId	int	   `json:"IdEmploye"`
}