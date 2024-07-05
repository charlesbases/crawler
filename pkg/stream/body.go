package stream

import (
	"bufio"
	"io"
	"sync"
)

// Line .
type Line []byte

// Body .
type Body interface {
	Read() []Line
	Write(dst io.Writer)
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
	r io.ReadCloser
}

// Response .
func Response(r io.ReadCloser) Body {
	return &response{r: r}
}

// Read .
func (r *response) Read() []Line {
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

	return lines
}

// backup to file
func (r *response) backup() {

}

// Save .
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
func (e extract) Read() []Line {
	return e
}

// Save .
func (e extract) Write(dst io.Writer) {
	for _, line := range e {
		dst.Write(line)
		dst.Write([]byte{'\n'})
	}
}
