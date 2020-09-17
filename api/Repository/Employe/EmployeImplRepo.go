/*
 * File Name EmployeImplRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package Repository

import (
	model "github.com/aditya37/backend-jobs/api/Model/Entity/Employe"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeImpl struct {
	Database *gorm.DB
}

func NewEmployeImpl (DBClient *gorm.DB) IEmployeRepo {
	return &EmployeImpl{Database: DBClient}
}

func (e *EmployeImpl) RegisterEmploye(addEmploye *model.EmployeAccount) *model.EmployeAccount {
	e.Database.Create(&addEmploye)
	return addEmploye
}

func (e *EmployeImpl) AddEmployeData(employeData *model.EmployeData) *model.EmployeData {
	e.Database.Create(&employeData)
	return employeData
}

func (e *EmployeImpl) AddEmployeAddress(employeAddr *model.EmployeAddress) *model.EmployeAddress {
	e.Database.Create(&employeAddr)
	return employeAddr
}

func (e *EmployeImpl) AddEmployeAttachment(employeAttach *model.EmployeAttachment) *model.EmployeAttachment {
	e.Database.Create(&employeAttach)
	return employeAttach
}

func (e *EmployeImpl) AddEmployeSocial(employeSocial *model.EmployeSocial) *model.EmployeSocial {
	e.Database.Create(&employeSocial)
	return employeSocial
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
