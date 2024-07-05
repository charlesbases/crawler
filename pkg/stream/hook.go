package stream

import "regexp"

// Hook .
type Hook func(r Body) Body

// FindSubmatch find sub string with regexp
func FindSubmatch(a string) Hook {
	x := regexp.MustCompile(a)

	return func(r Body) Body {
		var lines []Line
		for _, line := range r.Read() {
			if res := x.FindSubmatch(line); len(res) > 1 {
				lines = append(lines, res[1])
			}
		}
		return Extract(lines)
	}
}
