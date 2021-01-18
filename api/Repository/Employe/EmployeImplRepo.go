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
	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type EmployeImpl struct {
	Database *sqlx.DB
}

var (
	err error
)

func NewEmployeImpl(DBClient *sqlx.DB) IEmployeRepo {
	return &EmployeImpl{Database: DBClient}
}

// Function for hash or salting password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Function for decrypt password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (e *EmployeImpl) RegisterEmploye(addEmploye *model.EmployeAccount) (*model.EmployeAccount, error) {

	var err error

	TempData := []model.EmployeAccount{}
	err = e.Database.Select(&TempData, "SELECT username FROM employe_accounts WHERE username=$1", addEmploye.Username)
	if err != nil {
		return nil, err
	}

	if len(TempData) > 0 {
		return nil, errors.New("Username already exists")
	}

	// Hash password
	HashPassword, err := HashPassword(addEmploye.Password)
	addEmploye.Password = string(HashPassword)
	addEmploye.IsActive = "False"
	SQLInsert := `INSERT INTO employe_accounts(
		id,
		username,
		password,
		email,
		photo_profile,
		refresh_token,
		is_active) 
		VALUES(
			:id,
			:username,
			:password,
			:email,
			:photo_profile,
			:refresh_token,
			:is_active)`

	_, err = e.Database.NamedExec(SQLInsert, addEmploye)
	if err != nil {
		return nil, err
	}

	return addEmploye, nil
}

func (e *EmployeImpl) EmployeLogin(username, password string) ([]model.EmployeAccount, error) {

	var EmployeAccount []model.EmployeAccount
	err = e.Database.Select(&EmployeAccount, "SELECT username,id,password FROM employe_accounts WHERE username=$1 AND is_active='True'", username)
	if err != nil {
		return nil, err
	}

	if len(EmployeAccount) <= 0 {
		return nil, errors.New("Username not found")
	}
	// Verify password and dencrypt password
	err := VerifyPassword(EmployeAccount[0].Password, password)
	if err != nil {
		return nil, err
	}

	return EmployeAccount, nil
}

func (e *EmployeImpl) RefreshEmailVerify(email string) []model.EmployeAccount {
	var (
		Result []model.EmployeAccount
	)

	err := e.Database.Select(&Result, "SELECT * FROM employe_accounts WHERE email=$1 AND is_active='False'", email)
	if err != nil {
		return nil
	}
	return Result
}

func (e *EmployeImpl) EmployeEmailVerify(employeId string) error {

	TempData := []model.EmployeAccount{}
	err = e.Database.Select(&TempData, "SELECT id,is_active FROM employe_accounts WHERE id=$1 AND is_active='False'", employeId)
	if err != nil {
		return err
	}

	if len(TempData) <= 0 {
		return errors.New("Account Has Verified")
	}
	_, err := e.Database.Queryx("UPDATE employe_accounts SET is_active='True' WHERE id=$1", employeId)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeImpl) AddEmployeData(employeData *model.EmployeData) (*model.EmployeData, error) {

	TempData := []model.EmployeData{}
	err = e.Database.Select(&TempData, "SELECT employe_id FROM employe_data WHERE employe_id=$1", employeData.EmployeId)
	if err != nil {
		return nil, err
	}

	if len(TempData) >= 1 {
		return nil, errors.New("Duplicate Data")
	}

	SQLInsert := `INSERT INTO employe_data VALUES(
		:first_name,
		:last_name,
		:birth,
		:birth_place,
		:is_male,
		:phone,
		:about,
		:employe_id)`

	_, err := e.Database.NamedExec(SQLInsert, employeData)
	if err != nil {
		return nil, err
	}
	return employeData, nil
}

func (e *EmployeImpl) AddEmployeAddress(employeAddr *model.EmployeAddress) (*model.EmployeAddress, error) {

	TempData := []model.EmployeAddress{}
	err = e.Database.Select(&TempData, "SELECT employe_id FROM employe_addresses WHERE employe_id=$1", employeAddr.EmployeId)
	if err != nil {
		return nil, err
	}

	// Check data, if same data found
	if len(TempData) >= 1 {
		return nil, errors.New("Duplicate Data")
	}

	SQLInsert := `INSERT INTO employe_addresses VALUES(
		:country_name,
		:province_name,
		:district_name,
		:address_1,
		:address_2,
		:postal_code,
		:employe_id)`

	_, err := e.Database.NamedExec(SQLInsert, employeAddr)
	if err != nil {
		return nil, err
	}
	return employeAddr, nil
}

func (e *EmployeImpl) AddEmployeAttachment(employeAttach *model.EmployeAttachment) (*model.EmployeAttachment, error) {

	TempData := []model.EmployeAttachment{}
	err = e.Database.Select(&TempData, "SELECT employe_id FROM employe_attachments WHERE employe_id=$1", employeAttach.EmployeId)
	if err != nil {
		return nil, err
	}

	if len(TempData) > 1 {
		return nil, errors.New("Duplicate Data")
	}

	SQLInsert := `INSERT INTO employe_attachments VALUES(
		:portofolio_file,
		:resume_file,
		:resume_object,
		:portofolio_object,
		:employe_id)`

	_, err := e.Database.NamedExec(SQLInsert, employeAttach)
	if err != nil {
		return nil, err
	}
	return employeAttach, nil
}

func (e *EmployeImpl) AddEmployeSocial(employeSocial *model.EmployeSocial) (*model.EmployeSocial, error) {

	TempData := []model.EmployeSocial{}
	err = e.Database.Select(&TempData, "SELECT employe_id FROM employe_socials WHERE employe_id=$1", employeSocial.EmployeId)
	if err != nil {
		return nil, err
	}

	if len(TempData) >= 1 {
		return nil, errors.New("Duplicate Data")
	}

	SQLInsert := `INSERT INTO employe_socials VALUES(
		:portofolio_link,
		:github_link,
		:linkedin_link,
		:blog_link,
		:twitter_link,
		:employe_id)`

	_, err := e.Database.NamedExec(SQLInsert, employeSocial)
	if err != nil {
		return nil, err
	}

	return employeSocial, nil
}

func (e *EmployeImpl) AddEmployeExperience(experience *model.EmployeExperience) (*model.EmployeExperience, error) {

	sql := `INSERT INTO employe_experiences VALUES(
		:company_name,
		:job_title,
		:job_desc,
		:is_active,
		:start_work,
		:end_work,
		:employe_id)`

	_, err := e.Database.NamedExec(sql, experience)
	if err != nil {
		return nil, err
	}

	return experience, nil
}

func (e *EmployeImpl) AddEmployeEducation(employeEdu *model.EmployeEducation) (*model.EmployeEducation, error) {

	sql := `INSERT INTO employe_educations VALUES(
		:institution_name,
		:degree,
		:is_active,
		:start_education,
		:end_education,
		:employe_id)`

	_, err := e.Database.NamedExec(sql, employeEdu)
	if err != nil {
		return nil, err
	}
	return employeEdu, nil
}

func (e *EmployeImpl) GetEmployeById(employeId int) ([]model.EmployeAccount, error) {

	// FIXME: search how to handle sql.nullstring
	result := []model.EmployeAccount{}

	sql := `SELECT 
		employe_accounts.id,
		employe_accounts.username,
		employe_accounts.email,
		employe_accounts.photo_profile,
		employe_accounts.refresh_token,
		employe_accounts.is_active,
		employe_accounts.date_create,
		employe_accounts.date_update,
		COALESCE(employe_addresses.district_name,'NULL') AS district_name,
		COALESCE(employe_addresses.address_1,'NULL') AS address_1,
		COALESCE(employe_addresses.address_2,'NULL') AS address_2,
		COALESCE(employe_educations.institution_name,'NULL') AS institution_name,
		COALESCE(employe_educations.degree,'NULL') AS degree,
		COALESCE(
			TO_CHAR(
				employe_educations.start_education,'YYYY-MM-DD')
			,'NULL') AS start_education,
		COALESCE(
			TO_CHAR(
				employe_educations.end_education,'YYYY-MM-DD')
			,'NULL') AS end_education,
		COALESCE(employe_attachments.portofolio_file,'NULL') AS portofolio_file,
		COALESCE(employe_attachments.resume_file,'NULL') AS resume_file,
		COALESCE(employe_data.first_name,'NULL') AS first_name,
		COALESCE(employe_data.phone,'0') AS phone,
		COALESCE(employe_data.about,'NULL') AS about,
		COALESCE(employe_experiences.company_name,'') AS company_name,
		COALESCE(employe_experiences.job_title,'NULL') AS job_title,
		COALESCE(employe_experiences.job_desc,'NULL') AS job_desc,
		COALESCE(
			TO_CHAR(
				employe_experiences.start_work,'YYYY-MM-DD')
			,'NULL') AS start_work,
		COALESCE(
			TO_CHAR(
				employe_experiences.end_work,'YYYY-MM-DD'),
			'NULL') AS end_work,
		COALESCE(employe_socials.portofolio_link,'NULL') AS portofolio_link,
		COALESCE(employe_socials.github_link,'NULL') AS github_link,
		COALESCE(employe_socials.blog_link,'NULL') AS blog_link
		
	FROM 
   		employe_accounts
	LEFT JOIN
		employe_addresses ON employe_accounts.id = employe_addresses.employe_id
	LEFT JOIN
		employe_educations ON employe_accounts.id =  employe_educations.employe_id
	LEFT JOIN
		employe_attachments ON employe_accounts.id = employe_attachments.employe_id
	LEFT JOIN
		employe_data ON employe_accounts.id = employe_data.employe_id
	LEFT JOIN
		employe_experiences ON employe_accounts.id = employe_experiences.employe_id
	LEFT JOIN
		employe_socials ON employe_accounts.id = employe_socials.employe_id
	WHERE 
		employe_accounts.id=$1 AND employe_accounts.is_active='True'`

	_rows, err := e.Database.Queryx(sql, employeId)
	if err != nil {
		return nil, err
	}

	for _rows.Next() {

		var rows model.CombinedEmployeAccount

		// Do structScan
		if err := _rows.StructScan(&rows); err != nil {
			return nil, err
		}

		// Mapping dengan library carta
		if err := carta.Map(_rows.Rows, &result); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *EmployeImpl) DeleteAccount(employeId int) error {
	// e.Database.Delete(&model.EmployeAccount{},employeId)
	return nil
}
