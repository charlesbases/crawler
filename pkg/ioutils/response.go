package ioutils

import (
	"bufio"
	"io"
	"os"
	"sync"

	"github.com/charlesbases/crawler/internal/util"
)

// Line .
type Line []byte

// Reader .
type Reader interface {
	Read() ([]Line, error)
}

// Writer .
type Writer interface {
	Write(dst io.Writer)
}

// Body .
type Body interface {
	Reader
	Writer
}

var bufPool sync.Pool

// newBufReader .
func newBufReader(r io.Reader) *bufio.Reader {
	if v := bufPool.Get(); v != nil {
		buf := v.(*bufio.Reader)
		buf.Reset(r)
		return buf
	}
	return bufio.NewReader(r)
}

// response .
type response struct {
	f string
	r io.ReadCloser
}

// Response .
func Response(r io.ReadCloser) Body {
	return &response{r: r}
}

// Read .
func (r *response) Read() ([]Line, error) {
	// backup to file
	if len(r.f) != 0 {
		if err := r.backup(); err != nil {
			return nil, err
		}
	}

	defer r.r.Close()

	buf := newBufReader(r.r)

	var lines []Line

	for {
		if line, err := buf.ReadBytes('\n'); err != nil {
			break
		} else {
			lines = append(lines, line)
		}
	}

	return lines, nil
}

// backup .
func (r *response) backup() error {
	bkfile, err := util.CreateDir(r.f, true)
	if err != nil {
		return err
	}

	file, err := os.Create(bkfile)
	if err != nil {
		return err
	}

	// copy body to file
	io.Copy(file, r.r)
	r.r.Close()

	// new io
	file.Seek(0, 0)
	r.r = file

	return nil
}

// Write .
func (r *response) Write(dst io.Writer) {
	defer r.r.Close()

	io.Copy(dst, r.r)
}

// extract .
type extract []Line

// Extract .
func Extract(lines []Line) Body {
	return extract(lines)
}

// Read .
func (e extract) Read() ([]Line, error) {
	return e, nil
}

// Write .
func (e extract) Write(dst io.Writer) {
	for _, line := range e {
		dst.Write(line)
		dst.Write([]byte{'\n'})
	}
}
