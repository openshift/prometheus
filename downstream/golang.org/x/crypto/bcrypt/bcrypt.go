package bcrypt

import "errors"

func Cost([]byte) (int, error) {
	return 0, errors.New("not supported")
}

func CompareHashAndPassword(_ []byte, _ []byte) error {
	return errors.New("not supported")
}
