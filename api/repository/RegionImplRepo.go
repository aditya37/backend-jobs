/*
 * File Name RegionImplRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	"database/sql"

	model "github.com/aditya37/backend-jobs/api/model/Entity"
)

type RegionImpl struct {
	Psql *sql.DB
}

func NewRegionImpl(psqlClient *sql.DB) IRegionRepo {
	return &RegionImpl{Psql:psqlClient}
}

// Add District
func (p *RegionImpl) AddDistrict(district *model.District) (*model.District,error){
	begin,err := p.Psql.Begin()
	if err != nil {
		return nil,err
	}
	_,err = begin.Exec("INSERT INTO tbl_region VALUES($1,$2)",&district.IdDistrict,&district.DistrictName)
	if err != nil {
		return nil,err
	}
	begin.Commit()
	return district,err
}

// Add Country
func (p *RegionImpl) AddCountry(country *model.Country) (*model.Country,error){
	begin,err := p.Psql.Begin()
	if err != nil {
		return nil,err
	}
	_,err = begin.Exec("INSERT INTO tbl_country VALUES($1,$2)",&country.IdCountry,&country.CountryName)
	if err != nil {
		return nil,err
	}
	begin.Commit()
	return country,err
}

// Add Province
func (p *RegionImpl) AddProvince(province *model.Province) (*model.Province,error){
	begin,err := p.Psql.Begin()
	if err != nil {
		return nil,err
	}
	_,err = begin.Exec("INSERT INTO tbl_province VALUES($1,$2)",&province.IdProvince,&province.ProvinceName)
	if err != nil {
		return nil,err
	}
	begin.Commit()
	return province,err
}

// Get District 
func (p *RegionImpl) GetDistrict() ([]model.District,error){
	var (
		districts model.District
		arrDistrict []model.District
	)
	
	result,err := p.Psql.Query("SELECT * FROM tbl_district")
	if err != nil {
		return nil,err
	}

	for result.Next() {
		if err := result.Scan(&districts.IdDistrict,&districts.DistrictName); err != nil {
			return nil, err
		}
		arrDistrict = append(arrDistrict,districts)
	}
	return arrDistrict,nil
}

//  Get Country
func (p *RegionImpl) GetCountry() ([]model.Country,error) {
	var (
		countrys model.Country
		arrCountry []model.Country
	)

	result,err := p.Psql.Query("SELECT * FROM tbl_country")
	if err != nil {
		return nil,err
	}

	for result.Next() {
		if err := result.Scan(&countrys.IdCountry,&countrys.CountryName); err != nil {
			return nil,err
		}
		arrCountry = append(arrCountry,countrys)
	}
	return arrCountry,nil
}

// Get Province
func (p *RegionImpl) GetProvince()	([]model.Province,error) {
	var (
		provinces   model.Province
		arrProvince []model.Province
	)

	result,err := p.Psql.Query("SELECT * FROM tbl_province")
	if err != nil {
		return nil,err
	}

	for result.Next() {
		if err := result.Scan(&provinces.IdProvince,&provinces.ProvinceName); err != nil {
			return nil,err
		}
		arrProvince = append(arrProvince,provinces)
	}
	return arrProvince,nil
}