package fp

import (
	"errors"
)

var (
	ErrInvalidParam = errors.New("Invalid parameters")
)

func recover2err(a interface{}, def error) error {
	switch v := a.(type) {
	case nil:
		return nil
	case error:
		return v
	case string:
		return errors.New(v)
	default:
		return def
	}
}
