/*
 * File Name RedisConnImpl.go
 * Created on Mon Sep 21 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"context"
	"strconv"
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

func (r *RedisConnImpl) AddEmailVerify(id int64,token string) error {
	var (
		cntx = context.Background()
		err error
	)

	RedisClient := r.ConnectToRedis()
	
	// Convert Int64 to string
	IdToString := strconv.FormatInt(id,10)
	err = RedisClient.Set(cntx,IdToString,token,r.Exp*time.Second).Err()
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
func (r *RedisConnImpl) RemoveToken(key string) error {
	var (
		cntx = context.Background()
		err error
	)
	RedisClient := r.ConnectToRedis()
	err = RedisClient.Del(cntx,key).Err()
	if err != nil {
		return err
	}
	return nil
}