/*
 * File Name IRegionRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	RegionEntity "github.com/aditya37/backend-jobs/api/Model/Entity"
)
 
type IRegionRepo interface {
	AddDistrict(district *RegionEntity.District) *RegionEntity.District
	AddCountry(country 	 *RegionEntity.Country)  *RegionEntity.Country
	AddProvince(province *RegionEntity.Province) *RegionEntity.Province
	GetDistrict() []RegionEntity.District
	GetCountry()  []RegionEntity.Country
	GetProvince() []RegionEntity.Province
}