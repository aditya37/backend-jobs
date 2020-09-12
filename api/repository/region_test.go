/*
 * File Name repository_test.go
 * Created on Sat Sep 12 2020
 *
 * Copyright (c) 2020
 */

package repository_test

import (
	model "github.com/aditya37/backend-jobs/api/model/Entity"
	"github.com/aditya37/backend-jobs/api/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("RegionRepositoryTest", func() {
	var (
		repo repository.IRegionRepo
		err error
	)

	BeforeEach(func() {
		gdb,err := gorm.Open(postgres.Open("host=127.0.0.1 port=5432 user=admin dbname=db_jobs sslmode=disable password=lymousin"),&gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())
		repo = repository.NewRegionImpl(gdb)
	})

	Describe("RegionRepository",func() {
		// Entity Country --Start--
		Describe("RegionEntityCountry",func() {
			It("CountryIsEmpty",func(){
				GetCountries := repo.GetCountry()
				Expect(err).To(BeNil())
				Expect(len(GetCountries)).To(Equal(0))
			})
			
			It("AddCountry",func() {
				AddCountry := &model.Country{
					CountryName: "Indonesia",
				}
				InsertCountry := repo.AddCountry(AddCountry)
				Expect(InsertCountry.CountryName).To(Equal("Indonesia"))
			})

			It("CountryNotEmpty",func() {
				GetCountrys := repo.GetCountry()
				Expect(GetCountrys[0].CountryName).To(Equal("Indonesia"))
			})
		})
		// Entity Country --Finish--
		
		// Entity District --Start--
		Describe("RegionEntityDistrict",func() {
			It("DistrictIsEmpty",func() {
				GetDistricts := repo.GetDistrict()
				Expect(len(GetDistricts)).To(Equal(0))
			})

			It("AddDistrict",func() {
				district := &model.District{
					DistrictName: "Surabaya",
				}
				AddDistrict := repo.AddDistrict(district)
				Expect(AddDistrict.DistrictName).To(Equal("Surabaya"))
			})

			It("DistrictNotEmpty",func() {
				GetDistricts := repo.GetDistrict()
				Expect(GetDistricts[0].DistrictName).To(Equal("Surabaya"))
			})
		})
		// Entity District --Finish--
		
		// Entity Province --Start--
		Describe("RegionEntityProvince",func() {
			It("ProvinceIsEmpty",func() {
				GetProvinces := repo.GetProvince()
				Expect(len(GetProvinces)).To(Equal(0))
			})

			It("AddProvince",func() {
				province := &model.Province{
					ProvinceName: "Jawa Timur",
				}
				AddProvince := repo.AddProvince(province)
				Expect(AddProvince.ProvinceName).To(Equal("Jawa Timur"))
			})

			It("ProvinceNotEmpty",func() {
				GetProvinces := repo.GetProvince()
				Expect(GetProvinces[0].ProvinceName).To(Equal("Jawa Timur"))
			})
		})
		// Entity Province --Finish--
	})

	AfterEach(func() {
		
		Expect(err).ShouldNot(HaveOccurred())
	})

})
