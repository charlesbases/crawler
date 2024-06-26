package logger

import (
	"fmt"

	"github.com/charlesbases/colors"

	"github.com/charlesbases/crawler/internal/display"
)

// Debug .
func Debug(step int, a interface{}) {
	fmt.Println(display.Sprint(step, display.Tabulator, a))
}

// Debugf .
func Debugf(step int, format string, a ...interface{}) {
	fmt.Println(display.Sprintf(step, display.Tabulator, format, a...))
}

// Error .
func Error(a interface{}) {
	fmt.Println(colors.RedSprint(a))
}

// Errorf .
func Errorf(format string, a ...interface{}) {
	fmt.Println(colors.RedSprintf(format, a...))
}
