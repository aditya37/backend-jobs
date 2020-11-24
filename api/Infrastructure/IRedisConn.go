/*
 * File Name IRedisConn.go
 * Created on Mon Sep 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import model "github.com/aditya37/backend-jobs/api/Model"

type IRedisConn interface {
	RedisPing() (string,error)
	AddEmailVerify(id int64,token string) error
	VerifyEmail(key string) (string,error)
	RemoveToken(key string) error 
	// FIXME: Dead Code-1
	CreateAuth(userID int64,TokenDetail *model.TokenDetails) error
	// FIXME: Dead Code-2
	FetchAuth(auth *model.AccessDetails) (uint64,error)
	DeleteAuth(uuID string) (int64,error)
}