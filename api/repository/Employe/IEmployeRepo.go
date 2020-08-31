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
	RegisterEmploye(addEmploye *model.EployeAccount) (*model.EployeAccount,error)
	AddEmployeData(employeData *model.EmployeData) 	(*model.EmployeData,error)
	AddEmployeAddress(employeAddr *model.EmployeAddress) (*model.EmployeAddress,error)
	AddEmployeAttachment(employeAttach *model.EmployeAttachment) (*model.EmployeAttachment,error)
	AddEmployeSocial(employeSocial *model.EmployeSocial) (*model.EmployeSocial,error)
	AddEmployeExperience(experience *model.EmployeExperience)(*model.EmployeExperience,error)
}