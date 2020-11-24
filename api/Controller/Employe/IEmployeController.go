/*
 * File Name IEmployeController.go
 * Created on Sun Sep 27 2020
 *
 * Copyright (c) 2020
 */
package controller

import "github.com/labstack/echo/v4"

type IEmployeController interface {
	RegisterEmploye(c echo.Context) error
	LoginEmploye(c echo.Context) error
	GetEmployeById(c echo.Context) error
	VerifyEmail(c echo.Context) error
	RefreshEmailVerify(c echo.Context) error
	AddEmployeData(c echo.Context) error
	AddEmployeAddress(c echo.Context) error
	AddEmployeAttachment(c echo.Context) error
	AddEmployeEducation(c echo.Context) error
	AddEmployeExperience(c echo.Context) error
	AddEmployeSocial(c echo.Context) error
	EmployeLogOut(c echo.Context) error
	TestValidate(c echo.Context) (err error)
}