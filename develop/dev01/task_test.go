package main

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	got := getCurrentTime()
	diff := time.Since(got)
	if diff > 5*time.Second {
		t.Errorf("got %q, wanted %q", got, diff)
	}
}
