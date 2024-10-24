package hkdf

import (
	"bytes"
	"hash"
	"io"
)

func Expand(hash func() hash.Hash, pseudorandomKey, info []byte) io.Reader {
	return &bytes.Buffer{}
}
