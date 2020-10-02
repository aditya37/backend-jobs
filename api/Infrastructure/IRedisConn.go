/*
 * File Name IRedisConn.go
 * Created on Mon Sep 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

type IRedisConn interface {
	RedisPing() (string,error)
	AddEmailVerify(id int64,token string) error
	VerifyEmail(key string) (string,error)
	RemoveToken(key string) error 
}