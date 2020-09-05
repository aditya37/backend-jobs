/*
 * File Name EmployeImplRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	model "github.com/aditya37/backend-jobs/api/model/Entity/Employe"
	"gorm.io/gorm"
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
