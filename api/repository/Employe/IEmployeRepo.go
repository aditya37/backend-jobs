/*
 * File Name IEmployeRepo.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package repository

import (
	model "github.com/aditya37/backend-jobs/api/model/Entity/Employe"
)

type IEmployeRepo interface {
	RegisterEmploye(addEmploye *model.EmployeAccount) *model.EmployeAccount
	AddEmployeData(employeData *model.EmployeData) 	  *model.EmployeData
	AddEmployeAddress(employeAddr *model.EmployeAddress) *model.EmployeAddress
	AddEmployeAttachment(employeAttach *model.EmployeAttachment) *model.EmployeAttachment
	AddEmployeSocial(employeSocial *model.EmployeSocial) *model.EmployeSocial
	AddEmployeExperience(experience *model.EmployeExperience) *model.EmployeExperience
}