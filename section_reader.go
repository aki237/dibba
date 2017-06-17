package dibba

import (
	"io"
)

// SectionReader is a struct obeying the io.Reader interface
// used to read the section of a given io.ReadSeeker (dibba.Reader).
type SectionReader struct {
	db   *Reader
	nth  int64
	till int64
}

// Read is the method of SectionReader that makes it compatible with io.Reader interface.
// It reads only a section of a file specified in the unexported fields on the struct.
func (s *SectionReader) Read(p []byte) (int, error) {
	if s.nth >= s.till {
		return 0, io.EOF
	}
	_, err := s.db.box.Seek(s.nth, io.SeekStart)
	if err != nil {
		return 0, err
	}
	n, err := s.db.box.Read(p)
	if err != nil {
		return 0, err
	}
	s.nth += int64(n)
	if s.nth > s.till {
		p = p[:int64(n)-s.nth+s.till]
		_, err := s.db.box.Seek(s.till, io.SeekStart)
		if err != nil {
			return 0, err
		}
		return n - int(s.nth-s.till), nil
	}
	return n, nil
}
