package logger

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	var count = 5

	for i := 0; i < count; i++ {
		Debug(i, time.Now().Format(time.DateTime))
	}
	Error(count, time.Now().Format(time.DateTime))

	for i := 0; i < count; i++ {
		Debugf(i, "now: %s", time.Now().Format(time.DateTime))
	}
	Errorf(count, "now: %s", time.Now().Format(time.DateTime))
}
