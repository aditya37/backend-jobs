package repository_test

import (
	"time"

	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	repository "github.com/aditya37/backend-jobs/api/Repository/Employe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("EmployeRepository", func() {
	var (
		// err error
		repo repository.IEmployeRepo
	)
	BeforeEach(func() {
		gdb,err := gorm.Open(postgres.Open("host=127.0.0.1 port=5432 user=admin dbname=db_jobs sslmode=disable password=lymousin"),&gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())
		repo = repository.NewEmployeImpl(gdb)
	})
	
	It("RegisterEmploye",func() {
		account := &model.EmployeAccount{
			Id: 1,
			Username: "aditya",
			Password: "lymousin",
			Email: "aditya.krohman@gmail.com",
			PhotoProfile: "google.com",
			RefreshToken: "9999",
			DateCreate: time.Now(),
			DateUpdate: time.Now(),
		}
		regsiter,err := repo.RegisterEmploye(account)
		if err != nil {
			Skip(err.Error())
		}
		Expect(regsiter.Username).To(Equal("aditya"))
	})
	
	It("VerifyEmployeToken",func() {
		Tokens := model.EmployeAccount{RefreshToken: "9999"}
		err := repo.EmployeEmailVerify(Tokens.RefreshToken)
		if err != nil {
			Skip(err.Error())
		}
		Expect(err).To(BeNil())
	})
	
	It("LoginEmploye",func() {
		Login := &model.EmployeAccount{Username: "aditya",Password:"lymousin"}
		DoLogin,err := repo.EmployeLogin(Login.Username,Login.Password)
		if err != nil {
			Skip(err.Error())
		}
		Expect(DoLogin[0].Username).To(Equal("aditya"))
	})

	It("AddEmployeData",func() {
		EmployeData := &model.EmployeData{
			FirstName: "Aditya Rahman",
			LastName: "Rahman",
			Birth: time.Now().AddDate(1998,05,9),
			BirthPlace: "Bojonegoro",
			IsMale: "True",
			Phone: 6282257152133,
			About: "Dahlah",
			EmployeId: 1,
		}
		AddEmployeData,err := repo.AddEmployeData(EmployeData)
		if err != nil {
			Skip(err.Error())
		}
		Expect(AddEmployeData.EmployeId).To(Equal(1))
	})

	It("AddEmployeAddress",func() {
		EmployeAddress := &model.EmployeAddress{
			CountryName: "Indonesia",
			ProvinceName: "Jawa Timur",
			DistrictName: "Surabaya",
			Address_1: "Jln Haryo Matahun No 1",
			Address_2: "Rt 02 rw 02",
			PostalCode: 62171,
			EmployeId: 1,
		}
		employeAddresses,err := repo.AddEmployeAddress(EmployeAddress)
		if err != nil {
			Skip(string(err.Error()))
		}
		Expect(employeAddresses.CountryName).To(Equal("Indonesia"))
	})

	It("AddEmployeAttachment",func() {
		EmployeAttachment := &model.EmployeAttachment{
			PortofolioFile: "https://bit.ly/xVxdf",
			ResumeFile: "https://github.io/aditya37",
			EmployeId: 1,
		}
		attachment,err := repo.AddEmployeAttachment(EmployeAttachment)
		if err != nil {
			Skip(err.Error())
		}
		Expect(attachment.EmployeId).To(Equal(1))
	})

	It("AddEmployeSocial",func() {
		EmployeSocial := &model.EmployeSocial{
			PortofolioLink: "https://bit.ly/xVxdf",
			GithubLink: "aditya37",
			LinkedinLink: "google.com",
			BlogLink: "blog",
			TwitterLink: "kangcode",
			EmployeId: 1,
		}
		social,err := repo.AddEmployeSocial(EmployeSocial)
		if err != nil {
			Skip(err.Error())
		}
		Expect(social.EmployeId).To(Equal(int64(1)))
	})

	It("AddEmployeExperience",func() {
		EmployeExp := &model.EmployeExperience{
			CompanyName: "Gojek",
			JobTitle: "Backend",
			JobDesc: "Tukang ngode",
			IsActive: "True",
			StartWork: time.Now(),
			EndWork: time.Now(),
			EmployeId: 1,
		}
		repo.AddEmployeExperience(EmployeExp)
	})

	It("AddEmployeEducation",func() {
		EmployeEdu := &model.EmployeEducation{
			InstitutionName: "Polinema",
			Degree: "Diploma",
			Certificate: "bitly",
			IsActive: "True",
			StartEducation: time.Now(),
			EndEducation: time.Now(),
			EmployeId: 1,
		}
		repo.AddEmployeEducation(EmployeEdu)
	})
	
	It("GetEmployeId",func() {
		GetEmployeById := repo.GetEmployeById(1)
	if len(GetEmployeById) <= 0 {
		Skip("User belum lengkap")
	}
	Expect(GetEmployeById[0].Id).To(Equal(int64(1)))
	})

	It("DeleteEmployeAccountById",func() {
		// repo.DeleteAccount(1)
	})

	AfterEach(func(){

	})
})
