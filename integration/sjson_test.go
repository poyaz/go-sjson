package integration

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/poyaz/go-sjson/crypto"
	"github.com/poyaz/go-sjson/internal/codec"
)

type sjsonSuite struct {
	suite.Suite

	key []byte
}

func TestSjson_Encode_Decode(t *testing.T) {
	suite.Run(t, new(sjsonSuite))
}

func (s *sjsonSuite) SetupSuite() {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	randKey := func(length int, charset string) string {
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[seededRand.Intn(len(charset))]
		}
		return string(b)
	}(6, "abcdefghijklmnopqrstuvwxyz")

	s.key = []byte("00000000~" + randKey + "~00000000")
}

func (s *sjsonSuite) Test_sjson_encode_decode_int8() {
	enc, err := crypto.NewAesGcmEncryption(crypto.AesGcmOptions{EncryptionKey: s.key})
	s.NoError(err)

	codec, err := codec.NewJson(enc)
	s.NoError(err)

	_ = codec

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
