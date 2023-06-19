package crypto

import sjson "github.com/poyaz/go-sjson"

func GetDefaultEncryption(encryptionKey []byte) (sjson.Crypto, error) {
	return NewAesGcmEncryption(AesGcmOptions{EncryptionKey: encryptionKey})
}
