package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"time"
)

var secret string = os.Getenv("SECRET")

func getHash(time int) string {
	key := secret + "." + strconv.Itoa(time)
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
