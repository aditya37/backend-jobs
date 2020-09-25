/*
 * File Name RedisConnImpl.go
 * Created on Mon Sep 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConnImpl struct {
	RedisHost string
	Password string
	DB int
	Exp time.Duration
}

func NewRedisConn(host,password string, db int,exp time.Duration) IRedisConn {
	return &RedisConnImpl{
		RedisHost: host,
		Password: password,
		DB: db,
		Exp: exp,
	}
}

func (r *RedisConnImpl) ConnectToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr : r.RedisHost,
		Password:r.Password,
		DB:r.DB,
	})
}

func (r *RedisConnImpl) RedisPing() (string,error) {
	var (
		cntx = context.Background()
		err error
	)
	RedisClient := r.ConnectToRedis()
	pong,err := RedisClient.Ping(cntx).Result()
	return pong,err
}

func (r *RedisConnImpl) AddEmailVerify(email,token string) error {
	var (
		cntx = context.Background()
		err error
	)

	RedisClient := r.ConnectToRedis()
	err = RedisClient.Set(cntx,email,token,r.Exp*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisConnImpl) VerifyEmail(key string) (string,error) {
	var (
		cntx = context.Background()
		err error
	)

	RedisClient := r.ConnectToRedis()
	result,err := RedisClient.Get(cntx,key).Result()
	if err != nil {
		return "",err
	}
	return result,nil
}