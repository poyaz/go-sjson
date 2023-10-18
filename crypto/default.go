package crypto

import (
	sjson "github.com/poyaz/go-sjson/pkg"
)

func GetDefaultEncryption(encryptionKey []byte) (sjson.Crypto, error) {
	return NewAesGcmEncryption(AesGcmOptions{EncryptionKey: encryptionKey})
}
