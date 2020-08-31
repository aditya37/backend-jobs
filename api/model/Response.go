/*
 * File Name Response.go
 * Created on Thu Aug 27 2020
 *
 * Copyright (c) 2020
 */

package model

type SuccessResponse struct {
	Status 	int    `json:"status"`
	Message string `json:"message"`
	Result  interface{}
}