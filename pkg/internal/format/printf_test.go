package format

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintf(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(Sprint(i, Tabulator, time.Now()))
	}
	fmt.Println(Sprint(4, Space, time.Now()))
}
