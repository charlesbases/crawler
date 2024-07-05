package format

import (
	"fmt"
	"strings"
)

const (
	// Space .
	Space = "   "

	// Tabulator .
	Tabulator = "└─ "
)

// Sprint .
func Sprint(step int, prefix string, a interface{}) string {
	if step < 1 {
		return fmt.Sprint(a)
	}

	return fmt.Sprintf("%s%s%v", strings.Repeat("   ", step-1), prefix, a)
}

// Sprintf .
func Sprintf(step int, prefix string, format string, a ...interface{}) string {
	return Sprint(step, prefix, fmt.Sprintf(format, a...))
}
