// This package is a simple helper to
// read the database file as fixed sized
// pages.

package filepage

import (
	"errors"
	"io"
	"os"
)

const (
	maxPageBufferSize = 4096
)

// Error values
var (
	ErrorInvalidFile = errors.New("Invalid file")
)

// Scanner is a page reader
type Scanner struct {
	fd          *os.File
	fdSize      int64
	pageSize    int64
	currentPage int64
	pageBuffer  []byte
	err         error
}

// NewScanner returns new scanner for the specified file and pageSize
func NewScanner(filename string, pageSize int64) (*Scanner, error) {

	fileInfo, err := os.Stat(filename)

	if err != nil {
		return nil, err
	}

	if !fileInfo.Mode().IsRegular() {
		return nil, ErrorInvalidFile
	}

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return &Scanner{
		fd:          file,
		fdSize:      fileInfo.Size(),
		pageSize:    pageSize,
		currentPage: -1,
		pageBuffer:  make([]byte, maxPageBufferSize),
	}, nil
}

// Close release the opened file
func (s *Scanner) Close() {
	s.fd.Close()
}

// PagesNumber the number of pages
func (s *Scanner) PagesNumber() int64 {
	return s.fdSize / s.pageSize
}

// ReadPage reads the nect page into the buffer
func (s *Scanner) ReadPage() bool {

	var nextPage int64

	if s.currentPage >= 0 {
		nextPage = s.currentPage + 1
	}

	nextOffset := nextPage * s.pageSize

	if (nextOffset + s.pageSize) > s.fdSize {
		return false
	}
	_, err := s.fd.ReadAt(s.pageBuffer, nextOffset)

	if err != nil && err != io.EOF {
		s.err = err
		return false
	}

	s.currentPage = nextPage

	return true
}

// Page gives the current page value
func (s *Scanner) Page() []byte {
	cp := make([]byte, len(s.pageBuffer))
	copy(cp, s.pageBuffer)
	return cp
}

func (s *Scanner) CurrentOffset() int64 {
	return s.currentPage * s.pageSize
}

// Error returns the error occured during the reading
func (s *Scanner) Error() error {
	return s.err
}
