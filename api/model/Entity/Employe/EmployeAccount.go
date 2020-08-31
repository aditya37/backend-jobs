/*
 * File Name EmployeAccount.go
 * Created on Wed Aug 26 2020
 *
 * Copyright (c) 2020
 */

package model

import (
	"time"
)

type EployeAccount struct {
	Id	int		`json:"idEmploye"`
	Username string `json:"username"`
	Password string `json:"password"`
	RefreshToken string `json:"refreshToken"`
	DateCreate time.Time `json:"DateCreate"`
	DateUpdate time.Time `json:"DateUpdate"`
}