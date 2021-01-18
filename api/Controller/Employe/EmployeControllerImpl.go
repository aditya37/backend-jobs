/*
 * File Name EmployeControllerImpl.go
 * Created on Sun Sep 27 2020
 *
 * Copyright (c) 2020
 */

package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aditya37/backend-jobs/api/auth"
	util "github.com/aditya37/backend-jobs/api/utils"

	infrastructure "github.com/aditya37/backend-jobs/api/Infrastructure"
	response "github.com/aditya37/backend-jobs/api/Model"
	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	service "github.com/aditya37/backend-jobs/api/Service/Employe"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

type EmployeControllerImpl struct {
	EmployeService  service.IEmployeService
	EmployeRedis    infrastructure.IRedisConn
	FirebaseStorage infrastructure.IFireStorage
}

func NewEmployeController(s service.IEmployeService, r infrastructure.IRedisConn, fbs infrastructure.IFireStorage) IEmployeController {
	return &EmployeControllerImpl{
		EmployeService:  s,
		EmployeRedis:    r,
		FirebaseStorage: fbs,
	}
}

// FIXME: Ubah request dari json ke multipart form untuk handle file upload
func (e *EmployeControllerImpl) RegisterEmploye(c echo.Context) error {

	var err error

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	// Form untuk type file
	PhotoProfile, handler, err := c.Request().FormFile("photo_profile")

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// Form validate
	ValidatingRequest := model.EmployeAccount{
		Username: username,
		Password: password,
		Email:    email,
	}
	err = e.EmployeService.ValidateEmployeAccount(&ValidatingRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// File Handler and create temp file
	// FIXME: Buat function atau utils untuk membuat tempfile
	source, err := handler.Open()
	if err != nil {
		return err
	}
	defer source.Close()

	// Validate type file
	TypeFile, _ := util.GetContentType(source)
	log.Println(TypeFile)
	if TypeFile != "image/png" && TypeFile != "image/jpeg" {
		return c.JSON(http.StatusUnsupportedMediaType, response.SuccessResponse{
			Status:  0,
			Message: "File format not supported",
		})
	}

	// Create Temp File
	CreateTemp, err := util.CreateTemporyFile(os.TempDir(), "employe-photos-", PhotoProfile)
	if err != nil {
		log.Fatal(err)
	}

	// Upload to firebase
	UrlPhoto, err := e.FirebaseStorage.UploadPhotoProfile(username, CreateTemp.Name())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}
	os.Remove(CreateTemp.Name()) //cleaning up by removing the file

	// Save to database
	DoRegister, err := e.EmployeService.RegisterEmploye(&model.EmployeAccount{
		Username:     username,
		Password:     password,
		Email:        email,
		PhotoProfile: UrlPhoto,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// Generate new token
	Token := ksuid.New()
	err = e.EmployeRedis.AddEmailVerify(DoRegister.Id, Token.String())
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response.SuccessResponse{
			Status:  0,
			Message: "Opps,Server error",
		})
	}

	// Send email
	err = e.EmployeService.SendEmailVerify(DoRegister.Email, DoRegister.Username, Token.String(), DoRegister.Id)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.SuccessResponse{
			Status:  0,
			Message: "Failed to sent email",
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  1,
		Message: "Success, Please check your email inbox",
	})
}

func (e *EmployeControllerImpl) GetEmployeById(c echo.Context) error {

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	GetEmploye, err := e.EmployeService.GetEmployeById(EmployeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	if len(GetEmploye) < 1 {
		return c.JSON(http.StatusNotFound, response.SuccessResponse{
			Status:  0,
			Message: "Empty Data",
		})
	}
	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Success Load data",
		Result:  GetEmploye,
	})
}

func (e *EmployeControllerImpl) LoginEmploye(c echo.Context) error {
	var (
		EmployeAccount *model.EmployeAccount
	)

	// Decode json request
	err := json.NewDecoder(c.Request().Body).Decode(&EmployeAccount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// Do Login
	DoLogin, err := e.EmployeService.EmployeLogin(EmployeAccount.Username, EmployeAccount.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	GenerateToken, err := auth.GenerateToken(DoLogin[0].Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	if SaveToken := e.EmployeRedis.CreateAuth(DoLogin[0].Id, GenerateToken); SaveToken != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.SuccessResponse{
			Status:  0,
			Message: SaveToken.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Login Success",
		Result: map[string]interface{}{
			"idEmploye":     DoLogin[0].Id,
			"access_token":  GenerateToken.AccessToken,
			"refresh_token": GenerateToken.RefreshToken,
		},
	})
}

func (e *EmployeControllerImpl) VerifyEmail(c echo.Context) error {

	RedisKey := c.QueryParam("param")
	Token := c.QueryParam("token")

	DoVerify, err := e.EmployeRedis.VerifyEmail(RedisKey)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.SuccessResponse{
			Status:  0,
			Message: "Key not found",
		})
	}

	// Validate token from redis with query param
	if DoVerify != Token {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: "Wrong token",
		})
	}

	// Remove token/keys from redis
	e.EmployeRedis.RemoveToken(RedisKey)
	err = e.EmployeService.EmployeEmailVerify(RedisKey)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Account success registered",
	})
}

func (e *EmployeControllerImpl) RefreshEmailVerify(c echo.Context) error {

	var (
		err error
	)

	// FIXME: Ubah method ke post dan ganti formvalue
	EmailParam := c.FormValue("email")

	hasil := e.EmployeService.RefreshEmailVerify(EmailParam)
	log.Println(hasil)
	if len(hasil) <= 0 {
		return c.JSON(http.StatusNotFound, response.SuccessResponse{
			Status:  0,
			Message: "Oopps,Email not found",
		})
	}

	// Generate token
	Token := ksuid.New()
	err = e.EmployeRedis.AddEmailVerify(hasil[0].Id, Token.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// Send email
	err = e.EmployeService.SendEmailVerify(hasil[0].Email, hasil[0].Username, Token.String(), hasil[0].Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.SuccessResponse{
			Status:  0,
			Message: "Failed to sent email",
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  1,
		Message: "Refresh email verify success",
	})
}

func (e *EmployeControllerImpl) AddEmployeData(c echo.Context) error {

	var (
		EmployeData *model.EmployeData
	)

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	// Decode request
	err = json.NewDecoder(c.Request().Body).Decode(&EmployeData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// Validate user input
	err = e.EmployeService.ValidateEmployeData(EmployeData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	DumpData := model.EmployeData{
		FirstName:  EmployeData.FirstName,
		LastName:   EmployeData.LastName,
		Birth:      EmployeData.Birth, // FIXME: Atur Format untuk parsing time
		BirthPlace: EmployeData.BirthPlace,
		IsMale:     EmployeData.IsMale,
		Phone:      EmployeData.Phone,
		About:      EmployeData.About,
		EmployeId:  int64(EmployeId),
	}

	_, err = e.EmployeService.AddEmployeData(&DumpData)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  1,
		Message: "Success save data",
	})
}

func (e *EmployeControllerImpl) AddEmployeAddress(c echo.Context) error {

	var (
		EmployeAddress *model.EmployeAddress
	)

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	// Decode request
	err = json.NewDecoder(c.Request().Body).Decode(&EmployeAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	// Dump Data from request
	AddressData := model.EmployeAddress{
		CountryName:  EmployeAddress.CountryName,
		ProvinceName: EmployeAddress.ProvinceName,
		DistrictName: EmployeAddress.DistrictName,
		Address_1:    EmployeAddress.Address_1,
		Address_2:    EmployeAddress.Address_2,
		PostalCode:   EmployeAddress.PostalCode,
		EmployeId:    int64(EmployeId),
	}
	_, err = e.EmployeService.AddEmployeAddress(&AddressData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Success Add Address",
	})
}

// FIXME: Buat Menjadi pisah, tidak bisa di buat multiple ex: Portofolio upload
func (e *EmployeControllerImpl) AddEmployeAttachment(c echo.Context) error {

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	// Portofolio
	Portofolio, handler, err := c.Request().FormFile("portofolio_file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	PortofolioSource, err := handler.Open()
	if err != nil {
		return err
	}
	defer PortofolioSource.Close()

	// Validate type file
	PortofolioTypeFile, _ := util.GetContentType(PortofolioSource)
	if PortofolioTypeFile != "application/pdf" && PortofolioTypeFile != "application/msword" {
		return c.JSON(http.StatusUnsupportedMediaType, response.SuccessResponse{
			Status:  0,
			Message: "File format not supported",
		})
	}

	// Portofilio Temp
	PortofolioTemp, err := util.CreateTemporyFile(os.TempDir(), "employe-portofolio", Portofolio)
	if err != nil {
		return err
	}

	// Upload to Cloud Storage
	PortofolioURL, PortofolioObject, err := e.FirebaseStorage.UploadPortofolio(strconv.Itoa(EmployeId), PortofolioTemp.Name())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: "Failed to upload Portofolio",
		})
	}
	os.Remove(PortofolioTemp.Name())

	// Resume
	Resume, ResumeHandler, err := c.Request().FormFile("resume_file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	ResumeSource, err := ResumeHandler.Open()
	if err != nil {
		return err
	}
	defer ResumeSource.Close()

	// Validate type file
	ResumeTypeFile, _ := util.GetContentType(ResumeSource)
	if ResumeTypeFile != "application/pdf" && ResumeTypeFile != "application/msword" {
		return c.JSON(http.StatusUnsupportedMediaType, response.SuccessResponse{
			Status:  0,
			Message: "Resume File,format not supported",
		})
	}

	// Portofilio Temp
	ResumeTemp, err := util.CreateTemporyFile(os.TempDir(), "employe-resume", Resume)
	if err != nil {
		return err
	}

	// Upload to Cloud storage
	ResumeURL, ResumeName, err := e.FirebaseStorage.UploadResume(strconv.Itoa(EmployeId), ResumeTemp.Name())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: "Failed To Upload Resume",
		})
	}
	os.Remove(ResumeTemp.Name())

	// Dump Data
	data := &model.EmployeAttachment{
		PortofolioFile:   PortofolioURL,
		ResumeFile:       ResumeURL,
		ResumeObject:     ResumeName,
		PortofolioObject: PortofolioObject,
		EmployeId:        int64(EmployeId),
	}

	// Save to database
	_, err = e.EmployeService.AddEmployeAttachment(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Success Add Attachments",
	})
}

func (e *EmployeControllerImpl) TestValidate(c echo.Context) (err error) {
	return
}
func (e *EmployeControllerImpl) AddEmployeEducation(c echo.Context) error {

	var (
		EmployeEducation []*model.EmployeEducation
	)

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	err = json.NewDecoder(c.Request().Body).Decode(&EmployeEducation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	for i := 0; i < len(EmployeEducation); i++ {
		DumpData := model.EmployeEducation{
			InstitutionName: EmployeEducation[i].InstitutionName,
			Degree:          EmployeEducation[i].Degree,
			IsActive:        EmployeEducation[i].IsActive,
			StartEducation:  EmployeEducation[i].StartEducation,
			EndEducation:    EmployeEducation[i].EndEducation,
			EmployeId:       int64(EmployeId),
		}
		_, err = e.EmployeService.AddEmployeEducation(&DumpData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.SuccessResponse{
				Status:  0,
				Message: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Success add your educations",
	})
}

func (e *EmployeControllerImpl) AddEmployeExperience(c echo.Context) error {

	var (
		EmployeExperience []*model.EmployeExperience
	)

	// Validate Auth
	// Fixme: Bikin singleton biar gak buang2 rows
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	err = json.NewDecoder(c.Request().Body).Decode(&EmployeExperience)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	for i := 0; i < len(EmployeExperience); i++ {
		DumpData := model.EmployeExperience{
			CompanyName: EmployeExperience[i].CompanyName,
			JobTitle:    EmployeExperience[i].JobTitle,
			JobDesc:     EmployeExperience[i].JobDesc,
			IsActive:    EmployeExperience[i].IsActive,
			StartWork:   EmployeExperience[i].StartWork,
			EndWork:     EmployeExperience[i].EndWork,
			EmployeId:   int64(EmployeId),
		}

		_, err = e.EmployeService.AddEmployeExperience(&DumpData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.SuccessResponse{
				Status:  0,
				Message: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Success add your experience",
	})
}

func (e *EmployeControllerImpl) AddEmployeSocial(c echo.Context) error {

	var (
		EmployeSocial *model.EmployeSocial
	)

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}
	_, err = e.EmployeRedis.FetchAuth(ExtractJwtData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	EmployeId, _ := strconv.Atoi(c.Param("id"))
	err = json.NewDecoder(c.Request().Body).Decode(&EmployeSocial)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	DumpData := model.EmployeSocial{
		PortofolioLink: EmployeSocial.PortofolioLink,
		GithubLink:     EmployeSocial.GithubLink,
		LinkedinLink:   EmployeSocial.LinkedinLink,
		BlogLink:       EmployeSocial.BlogLink,
		TwitterLink:    EmployeSocial.TwitterLink,
		EmployeId:      int64(EmployeId),
	}

	_, err = e.EmployeService.AddEmployeSocial(&DumpData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.SuccessResponse{
			Status:  0,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, response.SuccessResponse{
		Status:  1,
		Message: "Success add your social media",
	})

}
func (e *EmployeControllerImpl) EmployeLogOut(c echo.Context) error {

	// Validate Auth
	ExtractJwtData, err := auth.ExtractMetaData(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	LogOut, err := e.EmployeRedis.DeleteAuth(ExtractJwtData.AccessUuid)
	if err != nil || LogOut == 0 {
		return c.JSON(http.StatusUnauthorized, response.SuccessResponse{
			Status:  0,
			Message: "unauthorized",
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  1,
		Message: "Successfully logged out",
	})
}
