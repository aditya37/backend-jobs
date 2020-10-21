/*
 * File Name IEchoRouter.go
 * Created on Sun Sep 27 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type IEchoRouter interface {
	StartServer(port string)
	RouterLogger()
	Get(uri string,f func(e echo.Context) error)
	Post(uri string,f func(e echo.Context) error)
	RouteGroup(uri string) *echo.Group
	ErrorHandler()
	CustomValidator(v *validator.Validate) 
}