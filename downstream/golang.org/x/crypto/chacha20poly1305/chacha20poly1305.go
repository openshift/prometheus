package chacha20poly1305

import (
	"crypto/cipher"
	"errors"
)

func New(key []byte) (cipher.AEAD, error) {
	return nil, errors.New("not supported")
}
