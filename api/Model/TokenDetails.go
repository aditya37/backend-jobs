/*
 * File Name TokenDetails.go
 * Created on Mon Nov 16 2020
 *
 * Copyright (c) 2020
 */

package model

type TokenDetails struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessUuid 	 string
	RefreshUuid  string
	AccessTokenExp int64
	RefreshTokenExp int64
}