package crawler

import (
	"testing"
)

func TestCrawler(t *testing.T) {
	c := New()

	c.Website("https://cn.bing.com").
		Save("./output/test/test")

	c.Run()
}
