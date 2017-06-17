package dibba

import "errors"

var (
	// ErrMalformed is returned if the given Reader is not a proper Dibba format file.
	ErrMalformed = errors.New("Malformed dibba passed")

	// ErrNoFileName is returned when no filename is specified for a given File struct.
	ErrNoFileName = errors.New("Filename not specified")

	// ErrFileAlreadyExists is returned while adding a new file to the package if a file
	// already exists in a package.
	ErrFileAlreadyExists = errors.New("File already exists")

	// ErrAlreadyCommitted is returned when Commit is called on a DibbaWriter struct again.
	ErrAlreadyCommitted = errors.New("Already committed")

	// ErrFileNotFound is returned when a given file is not found in the given package.
	ErrFileNotFound = errors.New("File not found in the Dibba")
)
