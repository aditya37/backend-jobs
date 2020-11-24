/*
 * File Name JwtMiddleware.go
 * Created on Thu Nov 19 2020
 *
 * Copyright (c) 2020
 */

// Fixme : Death Code
package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type (

	JwtMiddlewareConfig struct {
	
	// Berfungsi mengambil nilai jwt token dari request
	// Request yang sering digunakan
	// - "header:<Name>"
	// - "query:<Name>"
	// - "param:<Name>"
	// - "cookie:<Name>"
	ExtractToken string

	// AuthSchema digunakan di header Authorization
	// Default value "Bearer"
	AuthSchema string

	// Skipper middleware
	Skipper func(echo.Context) bool

	// Claims jwt.Claims

	}

	jwtExtractor func(echo.Context) (string, error)

)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)
var (
	DefaultJwtConfig = JwtMiddlewareConfig{
		Skipper: func(e echo.Context) bool {
			return false
		},
		ExtractToken: "header:" + echo.HeaderAuthorization,
		AuthSchema: "Bearer",
	}
)

func JwtMiddleware(config JwtMiddlewareConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultJwtConfig.Skipper
	}

	if config.ExtractToken == "" {
		config.ExtractToken = DefaultJwtConfig.ExtractToken
	}
	
	if config.AuthSchema == "" {
		config.AuthSchema = DefaultJwtConfig.AuthSchema
	}

	// Initalize
	parts := strings.Split(config.ExtractToken, ":")
	extractor := jwtFromHeader(parts[1],config.AuthSchema)
	switch parts[0] {
	case "query":
		extractor = JwtFormQuery(parts[1])
	case "param":
		extractor = JwtfromParam(parts[1])
	case "cookie":
		extractor = JwtFromCookie(parts[1])
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			auth,err := extractor(c)
			if err != nil {
				return err
			}

			VerifedToken,err := jwt.Parse(auth,func(token *jwt.Token) (interface{},error) {
				if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil,fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("SECRET_KEY")),nil
			})
			if err == nil && VerifedToken.Valid {
				return next(c)
			}
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid or expired jwt",
				Internal: err,
			}
		}
	}
}

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}

func JwtFormQuery(param string) jwtExtractor {
	return func(e echo.Context) (string, error) {
		AuthQuery := e.QueryParam(param)
		if AuthQuery == "" {
			return "",ErrJWTMissing
		}
		return AuthQuery,nil
	}
}

func JwtfromParam(param string) jwtExtractor {
	return func(e echo.Context) (string, error) {
		AuthParam := e.Param(param)
		if AuthParam == "" {
			return "",ErrJWTMissing
		}
		return AuthParam,nil
	}
}

func JwtFromCookie(name string) jwtExtractor {
	return func(e echo.Context) (string, error) {
		AuthCookie,err := e.Cookie(name)
		if err != nil {
			return "",ErrJWTMissing
		}
		return AuthCookie.Value,nil
	}
}