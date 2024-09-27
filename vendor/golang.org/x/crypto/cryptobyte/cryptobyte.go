package cryptobyte

import "errors"

type Builder struct{}

func (b *Builder) AddUint16(uint16) {
	return
}

func (b *Builder) AddUint8LengthPrefixed(func(*Builder)) {
}

func (b *Builder) AddBytes(v []byte) {
}

func (b *Builder) Bytes() ([]byte, error) {
	return nil, errors.New("not supported")
}
