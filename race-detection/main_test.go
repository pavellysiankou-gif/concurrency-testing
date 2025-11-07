package main

import "testing"

func TestCounter_Get(t *testing.T) {
	c := New(10)
	if got := c.Get(); got != 10 {
		t.Fatalf("Get() = %d, want 10", got)
	}
}
