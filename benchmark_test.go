//go:build bench
// +build bench

package go_sjson

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/poyaz/go-sjson/crypto"
	"github.com/poyaz/go-sjson/internal/codec"
	"github.com/poyaz/go-sjson/tests/helper"
)

func BenchmarkName(b *testing.B) {
	randKey := helper.GenerateStrRand(6)
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: []byte("00000000~" + randKey + "~00000000")})
	assert.NoError(b, err)

	codec, err := codec.NewJson(enc)
	assert.NoError(b, err)

	in := math.MaxUint32
	var out int

	for i := 0; i < b.N; i++ {
		mar, maErr := codec.Marshal(in)
		assert.NoError(b, maErr)

		unmErr := codec.Unmarshal(mar, &out)
		assert.NoError(b, unmErr)
	}
}
