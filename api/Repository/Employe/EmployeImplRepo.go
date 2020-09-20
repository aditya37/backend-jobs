/*
 * File Name EmployeImplRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	"errors"

	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeImpl struct {
	Database *gorm.DB
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
func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
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

func (e *EmployeImpl) EmployeEmailVerify(refreshToken string) error {
	
	TempData := &model.EmployeAccount{}
	if err = e.Database.Debug().Table("employe_accounts").Select("username,is_active").Where("refresh_token=?",refreshToken).Find(TempData).Error; err != nil {
		return err
	}

	// Check data, if token not valid or not registered
	if TempData.Username == "" {
		return errors.New("Data Not Found")
	}

	err = e.Database.Table("employe_accounts").Where("refresh_token=?",refreshToken).Update("is_active","True").Error
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeImpl) AddEmployeData(employeData *model.EmployeData) (*model.EmployeData,error){

	TempData := &model.EmployeData{}
	if err := e.Database.Table("employe_data").Select("employe_id").Where("employe_id=?",employeData.EmployeId).Find(TempData).Error; err != nil {
		return nil,err
	}
	
	// Check data, if same data found 
	if TempData.EmployeId != 0 {
		return nil,errors.New("Duplicate Employe Data")
	}

	e.Database.Create(&employeData)
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
	e.Database.Model(&model.EmployeAccount{}).Preload(clause.Associations).Where("employe_accounts.id=?",employeId).Find(&result)
	return result
}

func (e *EmployeImpl) DeleteAccount(employeId int) error {
	e.Database.Delete(&model.EmployeAccount{},employeId)
	return nil
}