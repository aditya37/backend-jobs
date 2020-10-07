/*
 * File Name jwtAuth.go
 * Created on Thu Oct 01 2020
 *
 * Copyright (c) 2020
 */

package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id int64) (string,error) {
	
	jwtClaims := jwt.MapClaims{}
	jwtClaims["authorized"] = true
	jwtClaims["id"] = id
	jwtClaims["exp"] = time.Now().Add(time.Hour*1).Unix()
	DoClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwtClaims)
	token,err := DoClaims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "",err
	}
	return token,nil
}
