/*
 * File Name IRegionRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	RegionEntity "github.com/aditya37/backend-jobs/api/model/Entity"
)
 
type IRegionRepo interface {
	AddDistrict(district *RegionEntity.District) (*RegionEntity.District,error)
	AddCountry(country 	 *RegionEntity.Country)  (*RegionEntity.Country,error)
	AddProvince(province *RegionEntity.Province) (*RegionEntity.Province,error)
	GetDistrict()	([]RegionEntity.District,error)
	GetCountry()	([]RegionEntity.Country,error)
	GetProvince()	([]RegionEntity.Province,error)	
}