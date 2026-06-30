package bac

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
)

// EncryptMessage encrypts a marshaled request JSON payload as BAC msg.
func EncryptMessage(plainJSON []byte, desKey string, nonce int64) (string, error) {
	if len(desKey) < 24 {
		return "", fmt.Errorf("bac: des key must be at least 24 bytes, got %d", len(desKey))
	}

	key := []byte(desKey[:24])
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("bac: create 3des cipher: %w", err)
	}

	padded, err := pkcs7Pad(plainJSON, block.BlockSize())
	if err != nil {
		return "", err
	}

	out := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, nonceIV(nonce))
	mode.CryptBlocks(out, padded)
	return base64.StdEncoding.EncodeToString(out), nil
}

func nonceIV(nonce int64) []byte {
	iv := make([]byte, des.BlockSize)
	binary.BigEndian.PutUint64(iv, uint64(nonce))
	return iv
}

func pkcs7Pad(src []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 || blockSize > 255 {
		return nil, fmt.Errorf("bac: invalid block size %d", blockSize)
	}
	padding := blockSize - len(src)%blockSize
	if padding == 0 {
		padding = blockSize
	}
	out := make([]byte, len(src)+padding)
	copy(out, src)
	for i := len(src); i < len(out); i++ {
		out[i] = byte(padding)
	}
	return out, nil
}

func pkcs7Unpad(src []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, fmt.Errorf("bac: invalid block size %d", blockSize)
	}
	if len(src) == 0 || len(src)%blockSize != 0 {
		return nil, errors.New("bac: invalid padded data length")
	}
	padding := int(src[len(src)-1])
	if padding == 0 || padding > blockSize || padding > len(src) {
		return nil, errors.New("bac: invalid padding")
	}
	for _, b := range src[len(src)-padding:] {
		if int(b) != padding {
			return nil, errors.New("bac: invalid padding")
		}
	}
	return src[:len(src)-padding], nil
}
