/*
 * File Name RegionImplService.go
 * Created on Sat Sep 05 2020
 *
 * Copyright (c) 2020
 */

package service

import (
	model "github.com/aditya37/backend-jobs/api/Model/Entity"
	repository "github.com/aditya37/backend-jobs/api/Repository"
)

type RegionService struct {
	RegionRepo repository.IRegionRepo
}


func NewRegionService(repository repository.IRegionRepo) IRegionService {
	return &RegionService{RegionRepo:repository}
}

func(r *RegionService) AddDistrict(district *model.District) *model.District {
	return r.RegionRepo.AddDistrict(district)
}

func(r *RegionService) AddCountry(country *model.Country) (*model.Country) {
	return r.RegionRepo.AddCountry(country)
}

func(r *RegionService) AddProvince(province *model.Province) *model.Province {
	return r.RegionRepo.AddProvince(province)
}

func(r *RegionService) GetDistrict() []model.District {
	return r.RegionRepo.GetDistrict()
}
	
func(r *RegionService) GetCountry() []model.Country {
	return r.RegionRepo.GetCountry()
}

func(r *RegionService) GetProvince() []model.Province {
	return r.RegionRepo.GetProvince()
}