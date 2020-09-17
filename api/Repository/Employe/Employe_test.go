package Repository_test

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
		RegisterEmploye := &model.EmployeAccount{
			Id: 2,
			Username: "Aditya",
			Password: "lymousin",
			PhotoProfile: "https://bit.ly/xnxxx",
			RefreshToken: "11111",
			DateCreate: time.Now(),
			DateUpdate: time.Now(),
		}
		insert := repo.RegisterEmploye(RegisterEmploye)
		Expect(insert.Username).To(Equal("Aditya"))
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
			EmployeId: 2,
		}
		AddEmployeData := repo.AddEmployeData(EmployeData)
		Expect(AddEmployeData.EmployeId).To(Equal(2))
	})

	It("AddEmployeAddress",func() {
		EmployeAddress := &model.EmployeAddress{
			CountryName: "Indonesia",
			ProvinceName: "Jawa Timur",
			DistrictName: "Surabaya",
			Address_1: "Jln Haryo Matahun No 1",
			Address_2: "Rt 02 rw 02",
			PostalCode: 62171,
			EmployeId: 2,
		}
		repo.AddEmployeAddress(EmployeAddress)
	})

	It("AddEmployeAttachment",func() {
		EmployeAttachment := &model.EmployeAttachment{
			PortofolioFile: "https://bit.ly/xVxdf",
			ResumeFile: "https://github.io/aditya37",
			EmployeId: 2,
		}
		repo.AddEmployeAttachment(EmployeAttachment)
	})

	It("AddEmployeSocial",func() {
		EmployeSocial := &model.EmployeSocial{
			PortofolioLink: "https://bit.ly/xVxdf",
			GithubLink: "adity37",
			LinkedinLink: "google.com",
			BlogLink: "blog",
			TwitterLink: "kangcode",
			EmployeId: 2,
		}
		repo.AddEmployeSocial(EmployeSocial)
	})

	It("AddEmployeExperience",func() {
		EmployeExp := &model.EmployeExperience{
			CompanyName: "Gojek",
			JobTitle: "Backend",
			JobDesc: "Tukang ngode",
			IsActive: "True",
			StartWork: time.Now(),
			EndWork: time.Now(),
			EmployeId: 2,
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
			EmployeId: 2,
		}
		repo.AddEmployeEducation(EmployeEdu)
	})

	// It("DeleteAccount",func() {
	// 	RegisterEmploye := &model.EmployeAccount{
	// 		Id: 1,
	// 		Username: "Aditya",
	// 		Password: "lymousin",
	// 		PhotoProfile: "https://bit.ly/xnxxx",
	// 		RefreshToken: "11111",
	// 		DateCreate: time.Now(),
	// 		DateUpdate: time.Now(),
	// 	}
	// 	repo.DeleteAccount(RegisterEmploye.Id)
	// })

	It("GetEmployeById",func() {
		
	})

	AfterEach(func(){

	})
})
