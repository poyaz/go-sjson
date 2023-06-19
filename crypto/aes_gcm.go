package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	sjson "github.com/poyaz/go-sjson"
)

type AesGcmEncryption struct {
	Cipher cipher.AEAD
	eType  sjson.EncryptType
}

type AesGcmOptions struct {
	// EncryptionKey is the encryption key used to encrypt the payloads
	// this key must be 16, 24, 32 characters in length
	EncryptionKey []byte
}

var _ sjson.Crypto = &AesGcmEncryption{}

func NewAesGcmEncryption(opts AesGcmOptions) (*AesGcmEncryption, error) {
	cipherBlock, err := aes.NewCipher(opts.EncryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	return &AesGcmEncryption{
		Cipher: gcm,
		eType:  AesGcm,
	}, nil
}

func (a *AesGcmEncryption) GetType() sjson.EncryptType {
	return a.eType
}

// Encrypt takes a byte array and returns an encrypted byte array
// as base64 encoded
func (a *AesGcmEncryption) Encrypt(unencryptedBytes []byte) ([]byte, error) {
	if len(unencryptedBytes) == 0 {
		return []byte(""), nil
	}

	nonce := make([]byte, a.Cipher.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedBytes := a.Cipher.Seal(nonce, nonce, unencryptedBytes, nil)
	encryptedEncodedData := make([]byte, base64.RawURLEncoding.EncodedLen(len(encryptedBytes)))
	base64.RawURLEncoding.Encode(encryptedEncodedData, encryptedBytes)

	return encryptedEncodedData, nil
}

// Decrypt takes an encrypted base64 byte array then
// returns an unencrypted byte array if same key was used to encrypt it
func (a *AesGcmEncryption) Decrypt(encryptedBytes []byte) ([]byte, error) {
	if len(encryptedBytes) == 0 {
		return []byte(""), nil
	}

	decodedEncryptedBytes := make([]byte, base64.RawURLEncoding.DecodedLen(len(encryptedBytes)))
	if _, err := base64.RawURLEncoding.Decode(decodedEncryptedBytes, encryptedBytes); err != nil {
		return nil, err
	}

	nonceSize := a.Cipher.NonceSize()
	if len(encryptedBytes) < nonceSize {
		return nil, NewKeyToShortError(len(encryptedBytes))
	}

	return a.Cipher.Open(
		nil,
		decodedEncryptedBytes[:nonceSize],
		decodedEncryptedBytes[nonceSize:],
		nil,
	)
}
