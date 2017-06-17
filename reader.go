package dibba

import (
	"encoding/binary"
	"io"
)

type fileBounds struct {
	filename string
	start    int64
	length   int64
}

type DibbaReader struct {
	box io.ReadSeeker
	fb  []fileBounds
}

// NewDibbaReader
func NewDibbaReader(rd io.ReadSeeker) *DibbaReader {
	return &DibbaReader{box: rd}
}

// Parse
func (db *DibbaReader) Parse() error {
	if err := db.checkIntegrity(); err != nil {
		return err
	}
	_, err := db.box.Seek(4, io.SeekStart)
	if err != nil {
		return err
	}
	for {
		nameheader := make([]byte, 2)
		n, err := db.box.Read(nameheader)
		if err != nil {
			if err != io.EOF || n != 1 {
				return err
			}
			if int(nameheader[0]) == TypeDibbaEnder {
				break
			}
		}
		if int(nameheader[0]) != TypeFile {
			break
		}
		fn := make([]byte, int(nameheader[1]))
		_, err = db.box.Read(fn)
		if err != nil {
			return err
		}
		var fileSize int64
		err = binary.Read(db.box, binary.BigEndian, &fileSize)
		if err != nil {
			return err
		}
		currentPos, err := db.box.Seek(fileSize, io.SeekCurrent)
		if err != nil {
			return err
		}
		nb := fileBounds{}
		nb.filename = string(fn)
		nb.start = currentPos - fileSize
		nb.length = fileSize
		db.fb = append(db.fb, nb)
	}
	return nil
}

func (db *DibbaReader) Open(filename string) (*File, error) {
	for _, val := range db.fb {
		if val.filename == filename {
			s := &SectionReader{db: db, nth: val.start, till: val.length + val.start}
			return &File{filename: filename, contents: s}, nil
		}
	}
	return nil, ErrFileNotFound
}

// checkIntegrity method returns boolean indicating whether the Dibba
// reader passed is consistent
func (db *DibbaReader) checkIntegrity() error {
	_, err := db.box.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	p := make([]byte, 4)
	_, err = db.box.Read(p)
	if err != nil {
		return err
	}
	if string(p) != string(DIBBA_HEADER) {
		return ErrMalformed
	}
	_, err = db.box.Seek(-1, io.SeekEnd)
	if err != nil {
		return err
	}
	p = make([]byte, 1)
	_, err = db.box.Read(p)
	if err != nil {
		return err
	}
	if string(p) != string(DIBBA_ENDER) {
		return ErrMalformed
	}
	return nil
}
