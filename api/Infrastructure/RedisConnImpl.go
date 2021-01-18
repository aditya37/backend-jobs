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

	model "github.com/aditya37/backend-jobs/api/Model"
	"github.com/go-redis/redis/v8"
)

type RedisConnImpl struct {
	RedisHost string
	Password  string
	DB        int
	Exp       time.Duration
}

func NewRedisConn(host, password string, db int, exp time.Duration) IRedisConn {
	return &RedisConnImpl{
		RedisHost: host,
		Password:  password,
		DB:        db,
		Exp:       exp,
	}
}

func (r *RedisConnImpl) ConnectToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.RedisHost,
		Password: r.Password,
		DB:       r.DB,
	})
}

func (r *RedisConnImpl) RedisPing() (string, error) {
	var (
		cntx = context.Background()
		err  error
	)
	RedisClient := r.ConnectToRedis()
	pong, err := RedisClient.Ping(cntx).Result()
	if err != nil {
		return "", err
	}
	return pong, nil
}

func (r *RedisConnImpl) AddEmailVerify(id int64, token string) error {
	var (
		cntx = context.Background()
		err  error
	)

	RedisClient := r.ConnectToRedis()

	// Convert Int64 to string
	IdToString := strconv.FormatInt(id, 10)
	err = RedisClient.Set(cntx, IdToString, token, r.Exp*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisConnImpl) VerifyEmail(key string) (string, error) {
	var (
		cntx = context.Background()
		err  error
	)

	RedisClient := r.ConnectToRedis()
	result, err := RedisClient.Get(cntx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func (r *RedisConnImpl) RemoveToken(key string) error {
	var (
		cntx = context.Background()
		err  error
	)
	RedisClient := r.ConnectToRedis()
	err = RedisClient.Del(cntx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

// FIXME: Dead Code-1
func (r *RedisConnImpl) CreateAuth(userID int64, TokenDetail *model.TokenDetails) error {

	var (
		cntx = context.Background()
	)

	AccessToken := time.Unix(TokenDetail.AccessTokenExp, 0)
	RefreshToken := time.Unix(TokenDetail.RefreshTokenExp, 0)
	now := time.Now()

	RedisClient := r.ConnectToRedis()
	AccesTokenErr := RedisClient.Set(cntx, TokenDetail.AccessUuid, strconv.Itoa(int(userID)), AccessToken.Sub(now)).Err()
	if AccesTokenErr != nil {
		return AccesTokenErr
	}

	RefreshTokenErr := RedisClient.Set(cntx, TokenDetail.RefreshUuid, strconv.Itoa(int(userID)), RefreshToken.Sub(now)).Err()
	if RefreshTokenErr != nil {
		return RefreshTokenErr
	}
	return nil
}

func (r *RedisConnImpl) FetchAuth(auth *model.AccessDetails) (uint64, error) {

	var (
		cntx = context.Background()
	)
	RedisClient := r.ConnectToRedis()

	AccessUuid, err := RedisClient.Get(cntx, auth.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	accId, _ := strconv.ParseUint(AccessUuid, 10, 64)
	return accId, nil
}

func (r *RedisConnImpl) DeleteAuth(uuID string) (int64, error) {

	var (
		cntx = context.Background()
	)
	RedisClient := r.ConnectToRedis()

	Deleted, err := RedisClient.Del(cntx, uuID).Result()
	if err != nil {
		return 0, err
	}

	return Deleted, nil
}
