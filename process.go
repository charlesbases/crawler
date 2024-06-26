package crawler

import (
	"os"

	"github.com/charlesbases/crawler/internal/util"
	"github.com/charlesbases/crawler/pkg/client"
	"github.com/charlesbases/crawler/pkg/ioutils"
)

// process .
type process interface {
	run(urls []ioutils.Line) ([]ioutils.Line, error)
}

// extractProcess .
type extractProcess struct {
	o *client.Options
	h ioutils.Handle
}

// newExtractProcess .
func newExtractProcess(o *client.Options, a ioutils.Handle) process {
	return &extractProcess{o: o, h: a}
}

// run .
func (p *extractProcess) run(urls []ioutils.Line) ([]ioutils.Line, error) {
	var lines = make([]ioutils.Line, 0, len(urls))
	for i := range urls {
		body, err := p.o.Run(urls[i])
		if err != nil {
			return nil, err
		}

		// handle with body
		abody, err := p.h(body)
		if err != nil {
			return nil, err
		}

		// read from abody
		alines, err := abody.Read()
		if err != nil {
			return nil, err
		}

		lines = append(lines, alines...)
	}
	return lines, nil
}

// outputProcess .
type outputProcess struct {
	o *client.Options
	f string
}

// newOutputProcess .
func newOutputProcess(o *client.Options, f string) process {
	return &outputProcess{o: o, f: f}
}

// run .
func (p *outputProcess) run(urls []ioutils.Line) ([]ioutils.Line, error) {
	absfile, err := util.CreateDir(p.f, false)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(absfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	for i := range urls {
		body, err := p.o.Run(urls[i])
		if err != nil {
			return nil, err
		}

		body.Write(file)
	}

	return urls, nil
}
