/*
 * File Name jwtAuth.go
 * Created on Thu Oct 01 2020
 *
 * Copyright (c) 2020
 */

package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	model "github.com/aditya37/backend-jobs/api/Model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GenerateToken(id int64) (*model.TokenDetails, error) {
	
	var (
		err error
	)
	
	DumpTokenDetail := &model.TokenDetails{}
	
	// Add Expired time for AccessToken
	DumpTokenDetail.AccessTokenExp = time.Now().Add(time.Minute * 15).Unix()
	// Create key for AccessToken to store in redis 
	DumpTokenDetail.AccessUuid = uuid.New().String()

	// Add Expire time for RefreshToken
	DumpTokenDetail.RefreshTokenExp = time.Now().Add(time.Hour * 24 * 7).Unix()
	// Create key for RefreshToken to store in redis
	DumpTokenDetail.RefreshUuid = uuid.New().String()

	// Creating access token
	jwtClaims := jwt.MapClaims{}
	jwtClaims["authorized"] = true
	jwtClaims["id"] = id
	jwtClaims["access_uuid"] = DumpTokenDetail.AccessUuid
	jwtClaims["exp"] = DumpTokenDetail.AccessTokenExp
	DoClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	
	DumpTokenDetail.AccessToken, err = DoClaims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// Creating refresh token
	RefreshTokenClaims := jwt.MapClaims{}
	RefreshTokenClaims["refresh_uuid"] = DumpTokenDetail.RefreshUuid
	RefreshTokenClaims["id"] = id
	RefreshTokenClaims["exp"] = DumpTokenDetail.RefreshTokenExp
	ClaimRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,RefreshTokenClaims)
	
	DumpTokenDetail.RefreshToken,err = ClaimRefreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return nil,err
	}
	
	return DumpTokenDetail, nil
}

func ExtractToken(c echo.Context) string {
	BearerToken := c.Request().Header.Get("Authorization")
	StrArr := strings.Split(BearerToken," ")
	if len(StrArr) == 2 {
		return StrArr[1]
	}
	return ""
}

func VerifyToken(c echo.Context) (*jwt.Token,error) {
	ExtractedToken := ExtractToken(c)
	token,err := jwt.Parse(ExtractedToken,func(token *jwt.Token) (interface{},error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")),nil
	})
	if err != nil {
		return nil,err
	}
	if _,ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil,err
	}
	return token,nil
}

func ExtractMetaData(c echo.Context) (*model.AccessDetails,error){
	VerifedToken,err := VerifyToken(c)
	if err != nil {
		return nil,err
	}
	claims,ok := VerifedToken.Claims.(jwt.MapClaims)
	if ok && VerifedToken.Valid {
		AccessUid,ok := claims["access_uuid"].(string)
		if !ok {
			return nil,err
		}
		return &model.AccessDetails{
			AccessUuid: AccessUid,
		},nil
	}
	return nil,nil
}