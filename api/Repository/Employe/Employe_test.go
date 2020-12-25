package repository_test

import (
	"strconv"
	"time"

	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	repository "github.com/aditya37/backend-jobs/api/Repository/Employe"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employe Repository", func() {
	
	var (
		EmployeRepo repository.IEmployeRepo
		// err error
	)
	
	BeforeEach(func() {
		
		db,err := sqlx.Connect("postgres","host=127.0.0.1 port=5432 user=admin dbname=db_jobs sslmode=disable password=lymousin")
		if err != nil {
			Expect(err).ShouldNot(HaveOccurred())
		}
		EmployeRepo = repository.NewEmployeImpl(db)
		
	})

	Describe("RegisterEmploye",func() {

		var (
			DataModel  model.EmployeAccount

			// DataModel is_active = "True"
			DataModel1 model.EmployeAccount
		)

		BeforeEach(func() {
			DataModel = model.EmployeAccount{
				Id:                7,
				Username:          "aditya98",
				Password:          "lymousin",
				Email:             "aditya.krohman@gmail.com",
				PhotoProfile:      "www.google.com",
				RefreshToken:      "7849",
				IsActive:          "False",
				DateCreate:        time.Now(),
				DateUpdate:        time.Now(),
			}

			DataModel1 = model.EmployeAccount{
				Id:                8,
				Username:          "danang98",
				Password:          "lymousin",
				Email:             "djaran_98@gmail.com",
				PhotoProfile:      "www.google.com",
				RefreshToken:      "7849",
				IsActive:          "True",
				DateCreate:        time.Now(),
				DateUpdate:        time.Now(),
			}

		})

		// Insert data into employe_accounts entity 
		It("Insert data into employe_accounts entity",func() {

			result,err := EmployeRepo.RegisterEmploye(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.Email).To(Equal("aditya.krohman@gmail.com"))

			result2,err := EmployeRepo.RegisterEmploye(&DataModel1)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result2.Email).To(Equal("djaran_98@gmail.com"))
		})

		// Insert data into employe_accounts entity if username exist
		It("Insert data into employe_accounts entity if username exist",func() {

			result,err := EmployeRepo.RegisterEmploye(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.Email).To(Equal("aditya.krohman@gmail.com"))

		})

		// Insert data into employe_accounts entity with username length 30
		It("Insert data into employe_accounts entity with username length 30", func() {

			DataModel.Username = "aditya9811aditya9811aditya98111"
			result,err := EmployeRepo.RegisterEmploye(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.Email).To(Equal("aditya.krohman@gmail.com"))

		})

	})

	Describe("EmployeEmailVerify",func() {
		
		var (
			DataModel model.EmployeAccount
			DataModel1 model.EmployeAccount
			err error
		)

		BeforeEach(func() {
			
			DataModel = model.EmployeAccount{
				Id:7,
				Email:"aditya.krohman@gmail.com",
			}

			DataModel1 = model.EmployeAccount{
				Id: 8,
				Email:"djaran_98@gmail.com",
			}

		})

		// Verify employe account with verifed employe_id
		It("Verify employe account with verifed employe_id",func() {
			
			// convert int64 to string
			strID := strconv.Itoa(int(DataModel1.Id))

			err = EmployeRepo.EmployeEmailVerify(strID)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}

		})
		
		// Verify employe account with unverifed employe_id 
		It("Verify employe account with unverifed employe_id", func() {
			
			// convert int64 to string
			strID := strconv.Itoa(int(DataModel.Id))

			err = EmployeRepo.EmployeEmailVerify(strID)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}

		})
		
		// Verify employe account with not registered email 
		It("Verify employe account with not registered email", func() {
			
			// Not registered account
			NotRegisteredModel := &model.EmployeAccount{
				Id: 99,
				Email:"guyon_waton@gmail.com",
			}
			
			strId := strconv.Itoa(int(NotRegisteredModel.Id))
			err = EmployeRepo.EmployeEmailVerify(strId)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}

		})

	})

	Describe("EmployeLogin",func() {
		
		var (
			DataModel model.EmployeAccount
			DataModel1 model.EmployeAccount
		)

		BeforeEach(func() {
			DataModel = model.EmployeAccount{
				Id:                7,
				Username:          "aditya98",
				Password:          "lymousin",
			}
			DataModel1 = model.EmployeAccount{
				Username:          "danang98",
				Password:          "lymousin",
			}
		})

		// Login with registered username
		It("Login with registered username",func ()  {

			result,err := EmployeRepo.EmployeLogin(DataModel1.Username,DataModel1.Password)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result[0].Username).To(Equal(DataModel1.Username))

		})

		// Login with username not verified
		It("Login with username not verified", func() {

			result,err := EmployeRepo.EmployeLogin(DataModel.Username,DataModel.Password)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result[0].Username).To(Equal(DataModel.Username))

		})

		// Login with wrong username and password
		It("Login with wrong username and password", func() {
			result,err := EmployeRepo.EmployeLogin("danang98","999999")
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result[0].Username).To(Equal(DataModel1.Username))
		})

	})

	Describe("AddEmployeData", func() {

		var (
			// Verifed account
			DataModel model.EmployeData
		)

		BeforeEach(func() {
			
			DataModel = model.EmployeData{
				FirstName:  "Danang",
				LastName:   "Imanto",
				Birth:      "1998-04-01",
				BirthPlace: "Dander",
				IsMale:     "True",
				Phone:      6282257152133,
				About:      "Ya Gitu",
				EmployeId:  8,
			}
		})

		// Add employe data with exist employe_id
		It("Add employe data with exist employe_id", func() {

			result,err := EmployeRepo.AddEmployeData(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(DataModel.EmployeId))

		})

		// Add employe data with not exist employe_id
		It("Add employe data with not exist employe_id", func() {

			result,err := EmployeRepo.AddEmployeData(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(DataModel.EmployeId))

		})
		
		// Add employe phone number with character ex: 0822-5715-2133
		It("Add employe phone number with character ex: 0822-5715-2133", func() {

			DataModel.Phone = 6822-5715-2133
			result,err := EmployeRepo.AddEmployeData(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(DataModel.EmployeId))

		})

	})

	Describe("AddEmployeAddress", func() {

		var (
			DataModel model.EmployeAddress
		)
		
		BeforeEach(func() {
			DataModel = model.EmployeAddress{
				CountryName:  "Indonesia",
				ProvinceName: "Jawa Timur",
				DistrictName: "Dander",
				Address_1:    "Jln Haryo",
				Address_2:    "Rt 01",
				PostalCode:   62171,
				EmployeId:    8,
			}
		})

		// Add employe address with exist employe_id
		It("Add employe address with exist employe_id", func() {
			result,err := EmployeRepo.AddEmployeAddress(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(DataModel.EmployeId))
		})

		// Add employe address with not exist employe_id
		It("Add employe address with not exist employe_id", func() {
			result,err := EmployeRepo.AddEmployeAddress(&DataModel)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(DataModel.EmployeId))
		})

	})

	Describe("AddEmployeAttachment", func() {

		var (
			ModelData  model.EmployeAttachment
		)
		
		BeforeEach(func() {
			ModelData = model.EmployeAttachment{
				PortofolioFile:   "www.google.com",
				ResumeFile:       "www.google.com",
				ResumeObject:     "resume",
				PortofolioObject: "portofolio",
				EmployeId:        8,
			}
		})

		// Add employe attachment with exist employe_id
		It("Add employe attachment with exist employe_id", func() {
			result,err := EmployeRepo.AddEmployeAttachment(&ModelData)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(ModelData.EmployeId))
		})

		// Add employe attachment with not exist employe_id
		It("Add employe attachment with not exist employe_id", func() {
			result,err := EmployeRepo.AddEmployeAttachment(&ModelData)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(ModelData.EmployeId))
		})
	})

	Describe("AddEmployeSocial", func() {

		var (
			ModelData model.EmployeSocial
		)
		
		BeforeEach(func() {
			ModelData = model.EmployeSocial{
				PortofolioLink: "github.com/aditya37",
				GithubLink:     "wwww",
				LinkedinLink:   "wwwww",
				BlogLink:       "wwwwww",
				TwitterLink:    "wwwwwwwwww",
				EmployeId:      8,
			}
		})

		// Add employe social media with exist employe_id
		It("Add employe social media with exist employe_id", func() {
			result,err := EmployeRepo.AddEmployeSocial(&ModelData)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(ModelData.EmployeId))
		})

		// Add employe social media with not exist employe_id
		It("Add employe social media with not exist employe_id", func() {
			result,err := EmployeRepo.AddEmployeSocial(&ModelData)
			if err != nil {
				Expect(err).ShouldNot(HaveOccurred())
			}
			Expect(result.EmployeId).To(Equal(ModelData.EmployeId))
		})

	})
})
