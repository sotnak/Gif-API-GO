package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func getHash(time int) string {
	key := Env.Secret + "." + strconv.Itoa(time)
	hash := sha256.New()

	hash.Write([]byte(key))

	hashRes := hex.EncodeToString(hash.Sum(nil))
	return hashRes
}

func check(hash string) bool {
	nowTime := time.Now().UTC()

	now := nowTime.Minute() + 60*(nowTime.Hour()+12*nowTime.Day())
	before := now - 1

	if hash == getHash(now) {
		return true
	}

	if hash == getHash(before) {
		return true
	}

	return false
}
