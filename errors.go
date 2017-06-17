package dibba

import "errors"

var (
	ErrMalformed         = errors.New("Malformed dibba passed")
	ErrNoFileName        = errors.New("Filename not specified")
	ErrFileAlreadyExists = errors.New("File already exists")
	ErrAlreadyCommitted  = errors.New("Already committed")
	ErrFileNotFound      = errors.New("File not found in the Dibba")
)
