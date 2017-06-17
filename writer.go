// Package dibba contains utility to interact with the dibba boxed documents
package dibba

import (
	"io"
)

// Unexported constants
var (
	DIBBA_HEADER = []byte{byte(TypeDibbaHeader), 'D', 'I', 'B'}
	DIBBA_ENDER  = []byte{byte(TypeDibbaEnder)}
)

// Exported Constants
const (
	TypeDibbaHeader int = iota
	TypeFile
	TypeDibbaEnder
)

// files is a collection of *File structs
type files []*File

// DibbaWriter is used to add (write) files in a dibba package
type DibbaWriter struct {
	files files
	box   io.WriteSeeker
	fresh bool
}

// NewDibba returns DibbaWriter struct with the passed WriteSeeker
// as the output file.
func NewDibba(ws io.WriteSeeker) *DibbaWriter {
	return &DibbaWriter{box: ws, fresh: true}
}

// Add method is used to add a File to the Dibba File system.
// Returns error if a file of the same name already exists.
func (db *DibbaWriter) Add(file *File) error {
	for _, val := range db.files {
		if file.Name() == val.Name() {
			return ErrFileAlreadyExists
		}
	}
	db.files = append(db.files, file)
	return nil
}

// Commit method is used to write all contents to a io.ReedSeeker (*os.File is compatible)
// including the header, contents and the ender.
func (db *DibbaWriter) Commit() error {
	if db.fresh != true {
		return ErrAlreadyCommitted
	}
	db.fresh = false
	_, err := db.box.Write(DIBBA_HEADER)
	if err != nil {
		return err
	}
	for _, val := range db.files {
		err := val.MarshalTo(db.box)
		if err != nil {
			return err
		}
	}

	_, err = db.box.Write(DIBBA_ENDER)
	return err
}
