package main

import (
	"math/rand"
	"time"
)

const (
	minNum = 2176782336  // 36^6
	maxNum = 78364164096 // 36^7
)

func assign_rand() int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Int63n(maxNum-minNum) + minNum
}

func err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
