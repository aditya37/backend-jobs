/*
 * File Name RegionImplRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	model "github.com/aditya37/backend-jobs/api/model/Entity"
	"gorm.io/gorm"
)

type RegionImpl struct {
	database *gorm.DB
}

func NewRegionImpl(dbClient *gorm.DB) IRegionRepo {
	return &RegionImpl{database:dbClient}
}

// Add District
func (p *RegionImpl) AddDistrict(district *model.District) *model.District {
	p.database.Create(&district)
	return district
}

// Add Country
func (p *RegionImpl) AddCountry(country *model.Country) *model.Country {
	p.database.Create(&country)
	return country
}

// Add Province
func (p *RegionImpl) AddProvince(province *model.Province) *model.Province {
	p.database.Create(&province)
	return province
}

// Get District 
func (p *RegionImpl) GetDistrict() []model.District {
	var districts []model.District
	p.database.Find(&districts)
	return districts
}

//  Get Country
func (p *RegionImpl) GetCountry() []model.Country {
	var countrys []model.Country
	p.database.Find(&countrys)
	return countrys
}

// Get Province
func (p *RegionImpl) GetProvince() []model.Province {
	var provinces []model.Province
	p.database.Find(&provinces)
	return provinces
}