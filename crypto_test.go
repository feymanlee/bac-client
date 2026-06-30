package bac

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/binary"
	"testing"
)

func TestEncryptMsgUsesFirst24BytesOfDESKeyBigEndianIVAndPKCS7(t *testing.T) {
	nonce := int64(1710000000123)
	got, err := EncryptMessage([]byte(`{"instanceId":"i-123"}`), "123456789012345678901234ignored", nonce)
	if err != nil {
		t.Fatalf("EncryptMessage() error = %v", err)
	}

	want := "hMFSpArB7ijjYIijY1647+jU/eegIglx"
	if got != want {
		t.Fatalf("EncryptMessage() = %q, want %q", got, want)
	}

	raw, err := base64.StdEncoding.DecodeString(got)
	if err != nil {
		t.Fatalf("ciphertext is not base64: %v", err)
	}
	if len(raw)%8 != 0 {
		t.Fatalf("ciphertext length = %d, want multiple of 8", len(raw))
	}
}

func TestNonceIVIsBigEndianUint64(t *testing.T) {
	got := nonceIV(0x0102030405060708)
	want := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	if !bytes.Equal(got, want) {
		t.Fatalf("nonceIV() = %v, want %v", got, want)
	}

	if binary.BigEndian.Uint64(got) != 0x0102030405060708 {
		t.Fatalf("nonceIV() is not big-endian")
	}
}

func TestEncryptMessageRejectsShortDESKey(t *testing.T) {
	_, err := EncryptMessage([]byte(`{}`), "short", 1)
	if err == nil {
		t.Fatal("EncryptMessage() error = nil, want error")
	}
}

func TestPKCS7PaddingAlwaysAddsBlock(t *testing.T) {
	got, err := pkcs7Pad([]byte("12345678"), 8)
	if err != nil {
		t.Fatalf("pkcs7Pad() error = %v", err)
	}
	if len(got) != 16 {
		t.Fatalf("padded length = %d, want 16", len(got))
	}
	for _, b := range got[8:] {
		if b != 8 {
			t.Fatalf("padding byte = %d, want 8", b)
		}
	}
}

func decrypt3DESCBCForTest(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	plain := make([]byte, len(ciphertext))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plain, ciphertext)
	return pkcs7Unpad(plain, block.BlockSize())
}
