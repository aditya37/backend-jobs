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
	EmployeService service.IEmployeService
	EmployeRedis infrastructure.IRedisConn
	FirebaseStorage infrastructure.IFireStorage
}

func NewEmployeController(s service.IEmployeService,r infrastructure.IRedisConn,fbs infrastructure.IFireStorage) IEmployeController {
	return &EmployeControllerImpl{
		EmployeService: s,
		EmployeRedis:r,
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
	PhotoProfile,handler,err := c.Request().FormFile("photo_profile")
	
	if err != nil {
		return c.JSON(http.StatusBadRequest,response.SuccessResponse{
			Status:0,
			Message:err.Error(),
		})
	}
	
	// Form validate
	ValidatingRequest := model.EmployeAccount{
		Username: username,
		Password: password,
		Email: email,
	}
	err = e.EmployeService.ValidateEmployeAccount(&ValidatingRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,response.SuccessResponse{
			Status: 0,
			Message: err.Error(),
		})
	}

	// File Handler and create temp file
	// FIXME: Buat function atau utils untuk membuat tempfile
	source,err := handler.Open()
	if err != nil {
		return err
	}
	defer source.Close()
	
	// Validate type file
	TypeFile,_ := util.GetContentType(source)
	if TypeFile != "image/jpeg" || TypeFile != "image/png" {
		return c.JSON(http.StatusUnsupportedMediaType,response.SuccessResponse{
			Status: 0,
			Message: "File format not supported",
		})
	}
	
	// Create Temp File
	CreateTemp,err := util.CreateTemporyFile(os.TempDir(),"employe-photos-",PhotoProfile)
	if err != nil {
		log.Fatal(err)
	}
	
	// Upload to firebase
	UrlPhoto,err := e.FirebaseStorage.UploadPhotoProfile(username,CreateTemp.Name())
	if err != nil {
		return c.JSON(http.StatusCreated,err.Error())
	}
	os.Remove(CreateTemp.Name()) //cleaning up by removing the file

	// Save to database
	DoRegister,err := e.EmployeService.RegisterEmploye(&model.EmployeAccount{
		Username: username,
		Password: password,
		Email: email,
		PhotoProfile: UrlPhoto,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError,response.SuccessResponse{
			Status: 0,
			Message: err.Error(),
		})
	}
	
	// Generate new token
	Token := ksuid.New()
	err = e.EmployeRedis.AddEmailVerify(DoRegister.Id,Token.String())
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,response.SuccessResponse{
			Status: 0,
			Message: "Opps,Server error",
		})
	}
	
	// Send email
	err = e.EmployeService.SendEmailVerify(DoRegister.Email,DoRegister.Username,Token.String(),DoRegister.Id)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError,response.SuccessResponse{
			Status: 0,
			Message: "Failed to sent email",
		})
	}

	return c.JSON(http.StatusCreated,response.SuccessResponse{
		Status: 1,
		Message: "Success, Please check your email inbox",
	})
}

func (e *EmployeControllerImpl) GetEmployeById(c echo.Context) error {
	
	EmployeId,_:= strconv.Atoi(c.Param("id"))
	GetEmploye := e.EmployeService.GetEmployeById(EmployeId)

	if len(GetEmploye) <= 0 {
		return c.JSON(http.StatusNotFound,response.SuccessResponse{
			Status:1,
			Message:"Data not found",
		})
	}
	return c.JSON(http.StatusAccepted,response.SuccessResponse{
		Status: 1,
		Message: "Success Load data",
		Result: GetEmploye,
	})
}

func (e *EmployeControllerImpl) LoginEmploye(c echo.Context) error {
	var (
		EmployeAccount *model.EmployeAccount
	)

	// Decode json request
	err := json.NewDecoder(c.Request().Body).Decode(&EmployeAccount)
	if err != nil {
		return c.JSON(http.StatusBadRequest,response.SuccessResponse{
			Status:0,
			Message:err.Error(),
		})
	}

	// Do Login 
	DoLogin,err := e.EmployeService.EmployeLogin(EmployeAccount.Username,EmployeAccount.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest,response.SuccessResponse{
			Status:0,
			Message:err.Error(),
		})
	}
	GenerateToken,err := auth.GenerateToken(DoLogin[0].Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,response.SuccessResponse{
			Status:0,
			Message:err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted,response.SuccessResponse{
		Status: 1,
		Message: "Login Success",
		Result:map[string]interface{}{
			"idEmploye" :DoLogin[0].Id,
			"access_token":string(GenerateToken),
		},
	})
}

func (e *EmployeControllerImpl) VerifyEmail(c echo.Context) error {
	
	RedisKey := c.QueryParam("param")
	Token 	 := c.QueryParam("token")

	DoVerify,err := e.EmployeRedis.VerifyEmail(RedisKey)
	if err != nil {
		return c.JSON(http.StatusNotFound,response.SuccessResponse{
			Status: 0,
			Message:"Key not found",
		})
	}
	
	// Validate token from redis with query param
	if DoVerify != Token {
		return c.JSON(http.StatusBadRequest,response.SuccessResponse{
			Status: 0,
			Message: "Wrong token",
		})
	}
	
	// Remove token/keys from redis
	e.EmployeRedis.RemoveToken(RedisKey)
	err = e.EmployeService.EmployeEmailVerify(RedisKey)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	return c.JSON(http.StatusAccepted,response.SuccessResponse{
		Status: 1,
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
	if len(hasil) <= 0 {
		return c.JSON(http.StatusNotFound,response.SuccessResponse{
			Status: 0,
			Message: "Oopps,Email not found",
		})
	}
	
	// Generate token
	Token := ksuid.New()
	err = e.EmployeRedis.AddEmailVerify(hasil[0].Id,Token.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,response.SuccessResponse{
			Status: 0,
			Message: err.Error(),
		})
	}
	
	// Send email
	err = e.EmployeService.SendEmailVerify(hasil[0].Email,hasil[0].Username,Token.String(),hasil[0].Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,response.SuccessResponse{
			Status: 0,
			Message: "Failed to sent email",
		})
	}

	return c.JSON(http.StatusCreated,response.SuccessResponse{
		Status: 1,
		Message: "Refresh email verify success",
	})
}