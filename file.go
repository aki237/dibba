package dibba

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type File struct {
	filename string
	contents io.Reader
}

func (f *File) GetReader() io.Reader {
	return f.contents
}

// Name method returns the filename of the File object
func (f *File) Name() string {
	return f.filename
}

// NewFile returns a File with the passed filename and Reader.
func NewFile(filename string, rd io.Reader) *File {
	return &File{filename: filename, contents: rd}
}

// marshalHeaders returns the bytes for the file header to be
// written in the Dibba file and an error if the filename is not set.
func (f *File) marshalHeaders() ([]byte, error) {
	if f.filename == "" {
		return nil, ErrNoFileName
	}
	b := bytes.NewBuffer(nil)
	_, err := b.Write([]byte{byte(TypeFile), byte(len(f.filename))})
	if err != nil {
		return nil, err
	}
	_, err = b.Write([]byte(f.filename))
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.BigEndian, int64(0))
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// MarshalTo method encodes the file contents (like in dibba file format)
// and writes it in a writer.
func (f *File) MarshalTo(w io.WriteSeeker) error {
	headers, err := f.marshalHeaders()
	if err != nil {
		return err
	}
	_, err = w.Write(headers)
	if err != nil {
		return err
	}
	pos, err := w.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}
	pos -= 8
	nWritten, err := io.Copy(w, f.contents)
	if err != nil {
		return err
	}
	_, err = w.Seek(pos, io.SeekStart)
	if err != nil {
		return err
	}
	fmt.Println(nWritten)
	err = binary.Write(w, binary.BigEndian, int64(nWritten))
	if err != nil {
		return err
	}
	_, err = w.Seek(0, io.SeekEnd)
	return err
}
