package logger

import (
	"fmt"

	"github.com/charlesbases/colors"

	"github.com/charlesbases/crawler/pkg/internal/format"
)

// Debug .
func Debug(step int, a interface{}) {
	fmt.Println(format.Sprint(step, format.Tabulator, a))
}

// Debugf .
func Debugf(step int, f string, a ...interface{}) {
	fmt.Println(format.Sprintf(step, format.Tabulator, f, a...))
}

// Error .
func Error(a interface{}) {
	fmt.Println(colors.RedSprint(a))
}

// Errorf .
func Errorf(f string, a ...interface{}) {
	fmt.Println(colors.RedSprintf(f, a...))
}
