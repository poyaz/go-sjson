//go:build integration
// +build integration

package integration

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/poyaz/go-sjson/crypto"
	sjson "github.com/poyaz/go-sjson/internal/codec"
	"github.com/poyaz/go-sjson/pkg"
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

	codec, err := sjson.NewJson(enc)
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

func (s *sjsonSuite) Test_sjson_encode_decode_map_int8() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := int8(10)
	maxInt8 := int8(math.MaxInt8)
	minInt8 := int8(math.MinInt8)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[int8]int8 data", func() {
			in := make(map[int8]int8)
			in[ten] = ten
			in[maxInt8] = maxInt8
			in[minInt8] = minInt8
			var out map[int8]int8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int8]int8 pointer data when deserialize", func() {
			in := make(map[int8]int8)
			in[ten] = ten
			in[maxInt8] = maxInt8
			in[minInt8] = minInt8
			var out *map[int8]int8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*int8]int8 data (pointer key)", func() {
			in := make(map[*int8]int8)
			in[&ten] = ten
			in[&maxInt8] = maxInt8
			in[&minInt8] = minInt8

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int8]*int8 data (pointer value)", func() {
			in := make(map[int8]*int8)
			in[ten] = &ten
			in[maxInt8] = &maxInt8
			in[minInt8] = &minInt8

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_int16() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := int16(10)
	maxInt16 := int16(math.MaxInt16)
	minInt16 := int16(math.MinInt16)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[int16]int16 data", func() {
			in := make(map[int16]int16)
			in[ten] = ten
			in[maxInt16] = maxInt16
			in[minInt16] = minInt16
			var out map[int16]int16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int16]int16 pointer data when deserialize", func() {
			in := make(map[int16]int16)
			in[ten] = ten
			in[maxInt16] = maxInt16
			in[minInt16] = minInt16
			var out *map[int16]int16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*int16]int16 data (pointer key)", func() {
			in := make(map[*int16]int16)
			in[&ten] = ten
			in[&maxInt16] = maxInt16
			in[&minInt16] = minInt16

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int16]*int16 data (pointer value)", func() {
			in := make(map[int16]*int16)
			in[ten] = &ten
			in[maxInt16] = &maxInt16
			in[minInt16] = &minInt16

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_int32() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := int32(10)
	maxInt32 := int32(math.MaxInt32)
	minInt32 := int32(math.MinInt32)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[int32]int32 data", func() {
			in := make(map[int32]int32)
			in[ten] = ten
			in[maxInt32] = maxInt32
			in[minInt32] = minInt32
			var out map[int32]int32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int32]int32 pointer data when deserialize", func() {
			in := make(map[int32]int32)
			in[ten] = ten
			in[maxInt32] = maxInt32
			in[minInt32] = minInt32
			var out *map[int32]int32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*int32]int32 data (pointer key)", func() {
			in := make(map[*int32]int32)
			in[&ten] = ten
			in[&maxInt32] = maxInt32
			in[&minInt32] = minInt32

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int32]*int32 data (pointer value)", func() {
			in := make(map[int32]*int32)
			in[ten] = &ten
			in[maxInt32] = &maxInt32
			in[minInt32] = &minInt32

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_int64() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := int64(10)
	maxInt64 := int64(math.MaxInt64)
	minInt64 := int64(math.MinInt64)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[int64]int64 data", func() {
			in := make(map[int64]int64)
			in[ten] = ten
			in[maxInt64] = maxInt64
			in[minInt64] = minInt64
			var out map[int64]int64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int64]int64 pointer data when deserialize", func() {
			in := make(map[int64]int64)
			in[ten] = ten
			in[maxInt64] = maxInt64
			in[minInt64] = minInt64
			var out *map[int64]int64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*int64]int64 data (pointer key)", func() {
			in := make(map[*int64]int64)
			in[&ten] = ten
			in[&maxInt64] = maxInt64
			in[&minInt64] = minInt64

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int64]*int64 data (pointer value)", func() {
			in := make(map[int64]*int64)
			in[ten] = &ten
			in[maxInt64] = &maxInt64
			in[minInt64] = &minInt64

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_int() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := 10
	maxInt := math.MaxInt
	minInt := math.MinInt
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[int]int data", func() {
			in := make(map[int]int)
			in[ten] = ten
			in[maxInt] = maxInt
			in[minInt] = minInt
			var out map[int]int

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int]int pointer data when deserialize", func() {
			in := make(map[int]int)
			in[ten] = ten
			in[maxInt] = maxInt
			in[minInt] = minInt
			var out *map[int]int

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*int]int data (pointer key)", func() {
			in := make(map[*int]int)
			in[&ten] = ten
			in[&maxInt] = maxInt
			in[&minInt] = minInt

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[int]*int data (pointer value)", func() {
			in := make(map[int]*int)
			in[ten] = &ten
			in[maxInt] = &maxInt
			in[minInt] = &minInt

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_uint8() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := uint8(10)
	maxUint8 := uint8(math.MaxUint8)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[uint8]uint8 data", func() {
			in := make(map[uint8]uint8)
			in[ten] = ten
			in[maxUint8] = maxUint8
			var out map[uint8]uint8

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint8]uint8 pointer data when deserialize", func() {
			in := make(map[uint8]uint8)
			in[ten] = ten
			in[maxUint8] = maxUint8
			var out *map[uint8]uint8

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*uint8]uint8 data (pointer key)", func() {
			in := make(map[*uint8]uint8)
			in[&ten] = ten
			in[&maxUint8] = maxUint8

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint8]*uint8 data (pointer value)", func() {
			in := make(map[uint8]*uint8)
			in[ten] = &ten
			in[maxUint8] = &maxUint8

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_uint16() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := uint16(10)
	maxUint16 := uint16(math.MaxUint16)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[uint16]uint16 data", func() {
			in := make(map[uint16]uint16)
			in[ten] = ten
			in[maxUint16] = maxUint16
			var out map[uint16]uint16

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint16]uint16 pointer data when deserialize", func() {
			in := make(map[uint16]uint16)
			in[ten] = ten
			in[maxUint16] = maxUint16
			var out *map[uint16]uint16

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*uint16]uint16 data (pointer key)", func() {
			in := make(map[*uint16]uint16)
			in[&ten] = ten
			in[&maxUint16] = maxUint16

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint16]*uint16 data (pointer value)", func() {
			in := make(map[uint16]*uint16)
			in[ten] = &ten
			in[maxUint16] = &maxUint16

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_uint32() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := uint32(10)
	maxUint32 := uint32(math.MaxUint32)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[uint32]uint32 data", func() {
			in := make(map[uint32]uint32)
			in[ten] = ten
			in[maxUint32] = maxUint32
			var out map[uint32]uint32

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint32]uint32 pointer data when deserialize", func() {
			in := make(map[uint32]uint32)
			in[ten] = ten
			in[maxUint32] = maxUint32
			var out *map[uint32]uint32

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*uint32]uint32 data (pointer key)", func() {
			in := make(map[*uint32]uint32)
			in[&ten] = ten
			in[&maxUint32] = maxUint32

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint32]*uint32 data (pointer value)", func() {
			in := make(map[uint32]*uint32)
			in[ten] = &ten
			in[maxUint32] = &maxUint32

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_uint64() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := uint64(10)
	maxUint64 := uint64(math.MaxUint64)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[uint64]uint64 data", func() {
			in := make(map[uint64]uint64)
			in[ten] = ten
			in[maxUint64] = maxUint64
			var out map[uint64]uint64

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint64]uint64 pointer data when deserialize", func() {
			in := make(map[uint64]uint64)
			in[ten] = ten
			in[maxUint64] = maxUint64
			var out *map[uint64]uint64

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*uint64]uint64 data (pointer key)", func() {
			in := make(map[*uint64]uint64)
			in[&ten] = ten
			in[&maxUint64] = maxUint64

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint64]*uint64 data (pointer value)", func() {
			in := make(map[uint64]*uint64)
			in[ten] = &ten
			in[maxUint64] = &maxUint64

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_uint() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	ten := uint(10)
	maxUint := uint(math.MaxUint)
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[uint]uint data", func() {
			in := make(map[uint]uint)
			in[ten] = ten
			in[maxUint] = maxUint
			var out map[uint]uint

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint]uint pointer data when deserialize", func() {
			in := make(map[uint]uint)
			in[ten] = ten
			in[maxUint] = maxUint
			var out *map[uint]uint

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*uint]uint data (pointer key)", func() {
			in := make(map[*uint]uint)
			in[&ten] = ten
			in[&maxUint] = maxUint

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[uint]*uint data (pointer value)", func() {
			in := make(map[uint]*uint)
			in[ten] = &ten
			in[maxUint] = &maxUint

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_string() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	str := "this-is-test"
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[string]string data", func() {
			in := make(map[string]string)
			in[str] = str
			var out map[string]string

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[string]string pointer data when deserialize", func() {
			in := make(map[string]string)
			in[str] = str
			var out *map[string]string

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*string]string data (pointer key)", func() {
			in := make(map[*string]string)
			in[&str] = str

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[string]*string data (pointer value)", func() {
			in := make(map[string]*string)
			in[str] = &str

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}

func (s *sjsonSuite) Test_sjson_encode_decode_map_bool() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := sjson.NewJson(enc)
	s.NoError(err)

	b := true
	var matcherNonTypeError *pkg.NonTypeError

	s.Run(
		"Should successfully encrypt/decrypt map[bool]bool data", func() {
			in := make(map[bool]bool)
			in[b] = b
			var out map[bool]bool

			mar, maErr := codec.Marshal(in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.NoError(unmErr)
			s.Equal(in, out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[bool]bool pointer data when deserialize", func() {
			in := make(map[bool]bool)
			in[b] = b
			var out *map[bool]bool

			mar, maErr := codec.Marshal(&in)
			unmErr := codec.Unmarshal(mar, &out)

			s.NoError(maErr)
			s.Error(unmErr)
			s.ErrorAs(unmErr, &matcherNonTypeError)
			s.Nil(out)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[*bool]bool data (pointer key)", func() {
			in := make(map[*bool]bool)
			in[&b] = b

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)

	s.Run(
		"Should error encrypt/decrypt map[bool]*bool data (pointer value)", func() {
			in := make(map[bool]*bool)
			in[b] = &b

			_, maErr := codec.Marshal(in)

			s.Error(maErr)
			s.ErrorAs(maErr, &matcherNonTypeError)
		},
	)
}
