/*
 * File Name EmployeImplService.go
 * Created on Thu Sep 17 2020
 *
 * Copyright (c) 2020
 */

package service

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"regexp"
	"time"

	util "github.com/aditya37/backend-jobs/api/utils"

	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	repository "github.com/aditya37/backend-jobs/api/Repository/Employe"
	"github.com/joho/godotenv"
)

type EmployeImplService struct {
	EmployeRepo repository.IEmployeRepo
}

func NewEmployeService(EmployeRepo repository.IEmployeRepo) IEmployeService {
	return &EmployeImplService{EmployeRepo: EmployeRepo}
}

func (e *EmployeImplService) ValidateEmployeAccount(employeAccount *model.EmployeAccount) error {
	
	if employeAccount == nil {
		err := errors.New("Empty employe account")
		return err
	}
	
	PasswordPattern,_ := regexp.Compile(`^[a-z0-9A-z\D]{8,60}$`)
	if len(employeAccount.Password) <= 8 {
		err := errors.New("Please use 8 character or more")
		return err
	}else if(PasswordPattern.MatchString(employeAccount.Password) != true ){
		err := errors.New("Please use strong password, you can use combined number or symbols")
		return err
	}

	EmailPattern,_ := regexp.Compile(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	isMatch := EmailPattern.MatchString(employeAccount.Email)
	if isMatch != true {
		err := errors.New("Wrong email format")
		return err
	}

	return nil
}

func (e *EmployeImplService) ValidateEmployeData(employeData *model.EmployeData) error {
	
	PhonePattern, _ := regexp.Compile(`^(^[1-9]|0)(\d{3,4}?){2}\d{3,4}$`)
	if PhonePattern.MatchString(string(employeData.Phone)) != true {
		err := errors.New("Please use this format 62xxx without plus (+)")
		return err
	}
	return nil
}

func (e *EmployeImplService) RegisterEmploye(addEmploye *model.EmployeAccount) (*model.EmployeAccount,error) {
	addEmploye.Id = time.Now().Unix()
	addEmploye.RefreshToken = util.StringWithCharSet(20)
	return e.EmployeRepo.RegisterEmploye(addEmploye)
}

func (e *EmployeImplService) EmployeLogin(username,password string) ([]model.EmployeAccount,error) {
	return e.EmployeRepo.EmployeLogin(username,password)
}

func (e *EmployeImplService) AddEmployeData(employeData *model.EmployeData) (*model.EmployeData,error) {
	return e.EmployeRepo.AddEmployeData(employeData)
}

func (e *EmployeImplService) AddEmployeAddress(employeAddr *model.EmployeAddress) (*model.EmployeAddress,error) {
	return e.EmployeRepo.AddEmployeAddress(employeAddr)
}

func (e *EmployeImplService) AddEmployeAttachment(employeAttach *model.EmployeAttachment) (*model.EmployeAttachment,error) {
	return e.EmployeRepo.AddEmployeAttachment(employeAttach)
}

func (e *EmployeImplService) AddEmployeSocial(employeSocial *model.EmployeSocial) (*model.EmployeSocial,error) {
	return e.EmployeRepo.AddEmployeSocial(employeSocial)
}

func (e *EmployeImplService) AddEmployeExperience(experience *model.EmployeExperience) *model.EmployeExperience {
	return e.EmployeRepo.AddEmployeExperience(experience)
}

func (e *EmployeImplService) AddEmployeEducation(employeEdu *model.EmployeEducation) *model.EmployeEducation {
	return e.EmployeRepo.AddEmployeEducation(employeEdu)
}

func (e *EmployeImplService) DeleteAccount(employeId int) error {
	return e.EmployeRepo.DeleteAccount(employeId)
}

func (e *EmployeImplService) GetEmployeById(employeId int) []model.EmployeAccount {
	return e.EmployeRepo.GetEmployeById(employeId)
}

func (e *EmployeImplService) EmployeEmailVerify(email string) error {
	return e.EmployeRepo.EmployeEmailVerify(email)
}

func (e *EmployeImplService) SendEmailVerify(to,employeUsername,employeToken string) error {
	
	// Load env variable for email authentication
	env := godotenv.Load()
	if env != nil {
		return nil
	}
	// Prepare and connection to email server
	auth := smtp.PlainAuth("",os.Getenv("SENDER"),os.Getenv("PASSWORD"),os.Getenv("HOST"))
	
	EmailTemplate,_ := template.ParseFiles("assets/pages/email_verify.html")
	var body bytes.Buffer

	MimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Don't Replay \n%s\n\n", MimeHeaders)))
	
	EmailTemplate.Execute(&body,struct {
		EmployeToken string
		EmployeUsername string
	}{
		EmployeToken: employeToken,
		EmployeUsername: employeUsername,
	})

	// Send email
	err  := smtp.SendMail(os.Getenv("ADDR"),auth,os.Getenv("SENDER"),[]string{to},body.Bytes())
	if err != nil {
		return err
	}

	return nil
}