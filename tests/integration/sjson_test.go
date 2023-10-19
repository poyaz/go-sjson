//go:build integration
// +build integration

package integration

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/poyaz/go-sjson/crypto"
	"github.com/poyaz/go-sjson/internal/codec"
	"github.com/poyaz/go-sjson/tests/helper"
)

type sjsonSuite struct {
	suite.Suite

	key []byte
}

func TestSjson_Encode_Decode(t *testing.T) {
	suite.Run(t, new(sjsonSuite))
}

func (s *sjsonSuite) SetupSuite() {
	s.key = []byte(helper.GenerateStrRand(16))
}

func (s *sjsonSuite) Test_sjson_encode_decode_int8() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt int8 data", func() {
			in := int8(10)
			var out int8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt int8 pointer data", func() {
			in := int8(10)
			var out int8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max int8 data", func() {
			in := int8(math.MaxInt8)
			var out int8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer int8 data", func() {
			in := int8(math.MaxInt8)
			var out int8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min int8 data", func() {
			in := int8(math.MinInt8)
			var out int8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min pointer int8 data", func() {
			in := int8(math.MinInt8)
			var out int8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_int16() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt int16 data", func() {
			in := int16(10)
			var out int16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt int16 pointer data", func() {
			in := int16(10)
			var out int16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max int16 data", func() {
			in := int16(math.MaxInt16)
			var out int16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer int16 data", func() {
			in := int16(math.MaxInt16)
			var out int16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min int16 data", func() {
			in := int16(math.MinInt16)
			var out int16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min pointer int16 data", func() {
			in := int16(math.MinInt16)
			var out int16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_int32() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt int32 data", func() {
			in := int32(10)
			var out int32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt int32 pointer data", func() {
			in := int32(10)
			var out int32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max int32 data", func() {
			in := int32(math.MaxInt32)
			var out int32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer int32 data", func() {
			in := int32(math.MaxInt32)
			var out int32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min int32 data", func() {
			in := int32(math.MinInt32)
			var out int32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min pointer int32 data", func() {
			in := int32(math.MinInt32)
			var out int32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_int64() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt int64 data", func() {
			in := int64(10)
			var out int64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt int64 pointer data", func() {
			in := int64(10)
			var out int64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max int64 data", func() {
			in := int64(math.MaxInt64)
			var out int64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer int64 data", func() {
			in := int64(math.MaxInt64)
			var out int64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min int64 data", func() {
			in := int64(math.MinInt64)
			var out int64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min pointer int64 data", func() {
			in := int64(math.MinInt64)
			var out int64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_int() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt int data", func() {
			in := 10
			var out int

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt int pointer data", func() {
			in := 10
			var out int

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max int data", func() {
			in := math.MaxInt
			var out int

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer int data", func() {
			in := math.MaxInt
			var out int

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min int data", func() {
			in := math.MinInt
			var out int

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt min pointer int data", func() {
			in := math.MinInt
			var out int

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_uint8() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt uint8 data", func() {
			in := uint8(10)
			var out uint8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt uint8 pointer data", func() {
			in := uint8(10)
			var out uint8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max uint8 data", func() {
			in := uint8(math.MaxUint8)
			var out uint8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer uint8 data", func() {
			in := uint8(math.MaxUint8)
			var out uint8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_uint16() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt uint16 data", func() {
			in := uint16(10)
			var out uint16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt uint16 pointer data", func() {
			in := uint16(10)
			var out uint16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max uint16 data", func() {
			in := uint16(math.MaxUint16)
			var out uint16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer uint16 data", func() {
			in := uint16(math.MaxUint16)
			var out uint16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_uint32() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt uint32 data", func() {
			in := uint32(10)
			var out uint32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt uint32 pointer data", func() {
			in := uint32(10)
			var out uint32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max uint32 data", func() {
			in := uint32(math.MaxUint32)
			var out uint32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer uint32 data", func() {
			in := uint32(math.MaxUint32)
			var out uint32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_uint64() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt uint64 data", func() {
			in := uint64(10)
			var out uint64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt uint64 pointer data", func() {
			in := uint64(10)
			var out uint64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max uint64 data", func() {
			in := uint64(math.MaxUint64)
			var out uint64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer uint64 data", func() {
			in := uint64(math.MaxUint64)
			var out uint64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_uint() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt uint data", func() {
			in := uint(10)
			var out uint

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt uint pointer data", func() {
			in := uint(10)
			var out uint

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max uint data", func() {
			in := uint(math.MaxUint)
			var out uint

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt max pointer uint data", func() {
			in := uint(math.MaxUint)
			var out uint

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_string() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt string data", func() {
			in := "this-is-test"
			var out string

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt uint string data", func() {
			in := "this-is-test"
			var out string

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_bool() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	s.Run(
		"Should successfully encrypt/decrypt bool data", func() {
			in := true
			var out bool

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should successfully encrypt/decrypt pointer bool data", func() {
			in := true
			var out bool

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)
}
