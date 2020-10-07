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
}