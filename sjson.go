package go_sjson

import (
	"github.com/poyaz/go-sjson/internal/codec"
	"github.com/poyaz/go-sjson/pkg"
)

func NewSjson(encryptionService pkg.Crypto) (codec.Json, error) {
	return codec.NewJson(encryptionService)
}
