package pkcs12

import (
	"encoding/pem"
	"errors"
)

func ToPEM(pfxData []byte, password string) ([]*pem.Block, error) {
	return nil, errors.New("not supported")
}
