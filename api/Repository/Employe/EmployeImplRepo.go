/*
 * File Name EmployeImplRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	"errors"
	"log"

	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeImpl struct {
	Database *gorm.DB
}

type QueryResult struct {
	IsActive string
	EmployeId int64
}

var (
	err error
)

func NewEmployeImpl (DBClient *gorm.DB) IEmployeRepo {
	return &EmployeImpl{Database: DBClient}
}

// Function for hash or salting password
func HashPassword(password string) ([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
}

// Function for decrypt password
func VerifyPassword(hashedPassword,password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}

func (e *EmployeImpl) RegisterEmploye(addEmploye *model.EmployeAccount) (*model.EmployeAccount,error) {
	
	var err error
	
	TempData := &model.EmployeAccount{}
	if err = e.Database.Table("employe_accounts").Select("username").Where("username=?",addEmploye.Username).Find(TempData).Error; err != nil {
		return nil,err
	}

	// Check data, if same data found 
	if TempData.Username != "" {
		return nil,errors.New("Registered Username, Please Use Another Username")
	}

	// Hash password
	HashPassword,err := HashPassword(addEmploye.Password)
	addEmploye.Password = string(HashPassword)
	addEmploye.IsActive = "False"
	e.Database.Create(&addEmploye)

	return addEmploye,nil
}

func (e *EmployeImpl) EmployeLogin(username,password string) ([]model.EmployeAccount,error) {
	
	var EmployeAccount []model.EmployeAccount

	if err := e.Database.Debug().Table("employe_accounts").Where("username=? AND is_active='True'",username).Take(&EmployeAccount).Error; err != nil {
		return nil,err
	}

	// Verify password and dencrypt password
	err := VerifyPassword(EmployeAccount[0].Password,password)
	if err != nil {
		return nil,err
	}

	return EmployeAccount,nil
}

func (e *EmployeImpl) RefreshEmailVerify(email string) []model.EmployeAccount {
	var (
		Result []model.EmployeAccount
	)
	e.Database.Model(&model.EmployeAccount{}).Select("*").Where("email=? AND is_active='False'",email).Find(&Result)
	return Result
}

func (e *EmployeImpl) EmployeEmailVerify(employeId string) error {
	
	TempData := &model.EmployeAccount{}
	if err = e.Database.Debug().Table("employe_accounts").Select("id,is_active='False'").Where("id=?",employeId).Find(TempData).Error; err != nil {
		return err
	}

	// Check data, if token not valid or not registered
	if TempData.Id == 0 {
		return errors.New("Account Not Found")
	}

	err = e.Database.Table("employe_accounts").Where("id=?",employeId).Update("is_active","True").Error
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeImpl) AddEmployeData(employeData *model.EmployeData) (*model.EmployeData,error){
	// Fixme : Buat verifikasi jika apakahh user sudah verifikasi
	TempData := &QueryResult{}
	
	if err := e.Database.Debug().Table("employe_accounts").Select("employe_accounts.is_active,employe_data.employe_id,employe_data.birth").Joins("inner join employe_data on employe_accounts.id=employe_data.employe_id").Where("employe_id=?",employeData.EmployeId).Find(&TempData).Error; err != nil {
		return nil,err
	}
	log.Println(TempData)
	if TempData.IsActive == "False" {
		return nil,errors.New("Account not verified")	
	}

	if TempData.EmployeId != 0 {
		return nil,errors.New("Duplicate data")
	}
	
	// Check data, if same data found 
	// if TempData.EmployeId != 0 {
	// 	return nil,errors.New("Duplicate Employe Data")
	// }

	if err := e.Database.Create(&employeData).Error; err != nil {
		return nil,err
	}
	return employeData,nil
}

func (e *EmployeImpl) AddEmployeAddress(employeAddr *model.EmployeAddress) (*model.EmployeAddress,error) {
	
	TempData := &model.EmployeAddress{}
	if err := e.Database.Table("employe_addresses").Select("employe_id").Where("employe_id=?",employeAddr.EmployeId).Find(TempData).Error; err != nil {
		return nil,err
	}
	
	// Check data, if same data found 
	if TempData.EmployeId != 0 {
		return nil,errors.New("Duplicate Address")
	}
	
	e.Database.Create(&employeAddr)
	return employeAddr,nil
}

func (e *EmployeImpl) AddEmployeAttachment(employeAttach *model.EmployeAttachment) (*model.EmployeAttachment,error) {
	
	TempData := &model.EmployeAttachment{}
	if err := e.Database.Table("employe_attachments").Select("employe_id").Where("employe_id=?",employeAttach.EmployeId).Find(TempData).Error; err != nil {
		return nil,err
	}
	
	// Check data, if same data found 
	if TempData.EmployeId != 0 {
		return nil,errors.New("Employe Attachment Has Been Filled")
	}
	
	e.Database.Create(&employeAttach)
	return employeAttach,nil
}

func (e *EmployeImpl) AddEmployeSocial(employeSocial *model.EmployeSocial) (*model.EmployeSocial,error) {
	
	TempData := &model.EmployeSocial{}
	if err := e.Database.Table("employe_socials").Select("employe_id=?",employeSocial.EmployeId).Find(TempData).Error; err != nil {
		return nil,err
	}
	
	// Check data, if same data found 
	if TempData.EmployeId != 0 {
		return nil,err
	}
	
	e.Database.Create(&employeSocial)
	return employeSocial,nil
}

func (e *EmployeImpl) AddEmployeExperience(experience *model.EmployeExperience) *model.EmployeExperience {
	e.Database.Create(&experience)
	return experience
}

func (e *EmployeImpl) AddEmployeEducation(employeEdu *model.EmployeEducation) *model.EmployeEducation {
	e.Database.Create(&employeEdu)
	return employeEdu
}

func (e *EmployeImpl) GetEmployeById(employeId int) []model.EmployeAccount{
	var result []model.EmployeAccount
	e.Database.Model(&model.EmployeAccount{}).Preload(clause.Associations).Where("employe_accounts.id=? AND is_active='True'",employeId).Find(&result)
	return result
}

func (e *EmployeImpl) DeleteAccount(employeId int) error {
	e.Database.Delete(&model.EmployeAccount{},employeId)
	return nil
}