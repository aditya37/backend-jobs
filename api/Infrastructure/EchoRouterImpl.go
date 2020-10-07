/*
 * File Name EchoRouterImpl.go
 * Created on Sun Sep 27 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoRouterImpl struct {
	EchoDispatcher *echo.Echo
}

func NewEchoRouter (EchoRouter *echo.Echo) IEchoRouter {
	return &EchoRouterImpl{EchoDispatcher: EchoRouter}
}

func (e *EchoRouterImpl) StartServer(port string) {
	e.EchoDispatcher.Logger.Fatal(e.EchoDispatcher.Start(port))
}

func (e *EchoRouterImpl) RouterLogger() {
	e.EchoDispatcher.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format : `=> ${time_rfc3339} ${status} ${method} ${host}${path} ${latency_name}` + "\n",
	}))
}

func (e *EchoRouterImpl) Get(uri string,f func(e echo.Context) error) {
	e.EchoDispatcher.GET(uri,f)
}

func (e *EchoRouterImpl) Post(uri string,f func(e echo.Context) error){
	e.EchoDispatcher.POST(uri,f)
}

// TODO: declare middleware skipper
func (e *EchoRouterImpl) RouteGroup(uri string) *echo.Group {
	return e.EchoDispatcher.Group(uri,middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			switch c.Request().RequestURI {
			case "/employes/signin":
				return true
			case "/employes/signup":
				return true
			default:
				break
			}
			return false
		},
		SigningMethod:"HS256",
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}))
}

// Handle error dari middleware
func (e *EchoRouterImpl) ErrorHandler() {
	e.EchoDispatcher.HTTPErrorHandler = func(e error, c echo.Context) {
		log.Println(e.Error())
		if http.StatusBadRequest == 400 {
			c.JSON(http.StatusBadRequest,echo.Map{
				"Status":0,
				"Message":"Bad Request",
			})
		}else if(http.StatusUnauthorized == 401){
			c.JSON(http.StatusUnauthorized,echo.Map{
				"Status":0,
				"Message":"Unauthorized",
			})
		}else{
			c.JSON(http.StatusBadGateway,echo.Map{
				"Status":0,
				"Message":e.Error(),
			})
		}
	}
}