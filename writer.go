/*Package dibba contains utility to interact with the dibba boxed documents.

Usage

This package can be used to interact with dibba file packages. All this operations
are done using structs and methods that are compatible with interfaces in io package
like "Reader", "Writer", "WriteSeeker" and etc.,

A sample usage is given below :

Consider inputFile is the dibbaFormat filename in the filesystem and file is the
filename of a file to be opened from the dibba package. Open the file.

  f, err := os.Open(inputFile)
  if err != nil {
  	fmt.Println(err)
  	return
  }

Create a new dibba.Reader using any kind of io.ReadSeeker.
In this case it is *os.File.

  d := dibba.NewReader(f)
  err = d.Parse()
  if err != nil {
  	fmt.Println(err)
  	return
  }

From the dibba.Reader open the file.

  fd, err := d.Open(file)
  if err != nil {
  	fmt.Println(err)
  	return
  }

Open returns a *dibba.File. There is a Reader inside the struct
which is obtained by GetReader method.

  _, err = io.Copy(os.Stdout, fd.GetReader())
  if err != nil {
  	fmt.Println(err)
  	return
  }

See the examples directory for handling dibba files using this package.
*/
package dibba

import (
	"io"
)

// Byte variables of Dibba identifier at start and end. Unexported.
var (
	dibbaHeader = []byte{byte(TypeDibbaHeader), 'D', 'I', 'B'}
	dibbaEnder  = []byte{byte(TypeDibbaEnder)}
)

// Exported Constants
const (
	TypeDibbaHeader int = iota
	TypeFile
	TypeDibbaEnder
)

// files is a collection of *File structs
type files []*File

// Writer is used to add (write) files in a dibba package
type Writer struct {
	files files
	box   io.WriteSeeker
	fresh bool
}

// NewWriter returns Writer struct with the passed WriteSeeker
// as the output file.
func NewWriter(ws io.WriteSeeker) *Writer {
	return &Writer{box: ws, fresh: true}
}

// Add method is used to add a File to the Dibba File system.
// Returns error if a file of the same name already exists.
func (db *Writer) Add(file *File) error {
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
func (db *Writer) Commit() error {
	if db.fresh != true {
		return ErrAlreadyCommitted
	}
	db.fresh = false
	_, err := db.box.Write(dibbaHeader)
	if err != nil {
		return err
	}
	for _, val := range db.files {
		err := val.MarshalTo(db.box)
		if err != nil {
			return err
		}
	}

	_, err = db.box.Write(dibbaEnder)
	return err
}
