package timeut

import (
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	Counter("test", func() {
		time.Sleep(time.Second)
	})
}
