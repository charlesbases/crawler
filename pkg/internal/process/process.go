package process

import (
	"os"

	"github.com/charlesbases/crawler/pkg/client"
	"github.com/charlesbases/crawler/pkg/internal/util"
	"github.com/charlesbases/crawler/pkg/stream"
)

// Process .
type Process interface {
	Run(urls []stream.Line) ([]stream.Line, error)
}

// extract .
type extract struct {
	o *client.Options
	h stream.Hook
}

// Extract .
func Extract(o *client.Options, hook stream.Hook) Process {
	return &extract{o: o, h: hook}
}

// Run .
func (e *extract) Run(urls []stream.Line) ([]stream.Line, error) {
	var lines = make([]stream.Line, 0, len(urls))
	for i := range urls {
		body, err := e.o.Run(urls[i])
		if err != nil {
			return nil, err
		}

		lines = append(lines, e.h(body).Read()...)
	}
	return lines, nil
}

// output .
type output struct {
	o *client.Options
	f string
}

// Output .
func Output(o *client.Options, file string) Process {
	return &output{o: o, f: file}
}

// Run .
func (o *output) Run(urls []stream.Line) ([]stream.Line, error) {
	absfile, err := util.CreateDirIfNotExist(o.f)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(absfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	for i := range urls {
		body, err := o.o.Run(urls[i])
		if err != nil {
			return nil, err
		}

		body.Write(file)
	}

	return urls, nil
}
