/*
 * File Name RandomToken.go
 * Created on Tue Sep 22 2020
 *
 * Copyright (c) 2020
 */

package util

import (
	"math/rand"
	"time"
)

const Charset = "abcdefghijklmnopqrstuvexyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var SeededRand *rand.Rand = rand.New(rand.NewSource(
	time.Now().UnixNano(),
))

func StringWithCharSet(length int) string {
	b := make([]byte,length)
	for i := range b {
		b[i] = Charset[SeededRand.Intn(len(Charset))]
	}
	return string(b)
}

func GenerateRefreshToken(length int) string {
	return StringWithCharSet(length)
}