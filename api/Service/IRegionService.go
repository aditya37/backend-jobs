/*
 * File Name IRegionService.go
 * Created on Sat Sep 05 2020
 *
 * Copyright (c) 2020
 */

package service

import (
	model "github.com/aditya37/backend-jobs/api/Model/Entity"
)

type IRegionService interface {
	AddDistrict(district *model.District) *model.District
	AddCountry(country 	 *model.Country)  (*model.Country)
	AddProvince(province *model.Province) *model.Province
	GetDistrict() []model.District
	GetCountry()  []model.Country
	GetProvince() []model.Province
}