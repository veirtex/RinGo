package errs

import "errors"

var (
	ErrArgs           = errors.New("args are not enough")
	ErrNotFound       = errors.New("key doesn't exist")
	ErrKeyDeleted     = errors.New("key deleted")
	ErrNilKey         = errors.New("key is nil")
	ErrDatatype       = errors.New("datatype is unknown")
	ErrUnknownCommand = errors.New("unknown command")
)
