package crawler

// Option .
type Option interface {
	apply(c *Crawler)
}

// optfn .
type optfn func(c *Crawler) ()

// apply .
func (fn optfn) apply(c *Crawler) () {
	fn(c)
}

// Concurrent .
func Concurrent(num int) Option {
	return optfn(func(c *Crawler) {
		c.concurrent = num
	})
}
