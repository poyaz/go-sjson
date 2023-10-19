//go:build integration || bench
// +build integration bench

package helper

import (
	"math/rand"
	"time"
)

const (
	_strRand = "abcdefghijklmnopqrstuvwxyz"
)

type GenerateStrRandOpts func(o GenerateStrRandOptions) GenerateStrRandOptions

type GenerateStrRandOptions struct {
	charset string
}

func GenerateStrRand(length int, opts ...GenerateStrRandOpts) string {
	options := GenerateStrRandOptions{charset: _strRand}
	for _, opt := range opts {
		options = opt(options)
	}

	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = options.charset[seededRand.Intn(len(options.charset))]
	}
	return string(b)
}
