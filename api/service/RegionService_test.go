/*
 * File Name RegionServiceTest.go
 * Created on Sun Sep 06 2020
 *
 * Copyright (c) 2020
 */

package service

import (
	"testing"

	model "github.com/aditya37/backend-jobs/api/model/Entity"
	"github.com/stretchr/testify/mock"
)

type RegionMockRepository struct {
	mock.Mock
}

func (m *RegionMockRepository) AddDistrict(district *model.District) *model.District {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.District)
}
func (m *RegionMockRepository) AddCountry(country *model.Country) *model.Country {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Country)
}

func (m *RegionMockRepository) AddProvince(province *model.Province) *model.Province {
	args := m.Called()
	result := args.Get(0)
	return result.(*model.Province)
}

func (m *RegionMockRepository) GetDistrict() []model.District {
	args := m.Called()
	result := args.Get(0)
	return result.([]model.District)
}

func (m *RegionMockRepository) GetCountry()  []model.Country {
	args := m.Called()
	result := args.Get(0)
	return result.([]model.Country)
}

func (m *RegionMockRepository) GetProvince() []model.Province {
	args := m.Called()
	result := args.Get(0)
	return result.([]model.Province)
}

func TestAddDistrict(t *testing.T){

}
