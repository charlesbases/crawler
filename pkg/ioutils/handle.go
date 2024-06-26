package ioutils

import "regexp"

// Handle .
type Handle func(r Body) (Body, error)

// FindSubmatch match with regexp
func FindSubmatch(a string) Handle {
	c := regexp.MustCompile(a)

	return func(r Body) (Body, error) {
		res, err := r.Read()
		if err != nil {
			return r, err
		}

		lines := make([]Line, 0, len(res))

		for _, line := range res {
			if v := c.FindSubmatch(line); len(v) != 0 {
				lines = append(lines, v[0])
			}
		}

		return Extract(lines), nil
	}
}
