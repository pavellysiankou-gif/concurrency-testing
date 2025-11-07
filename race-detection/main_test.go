package main

import "testing"

func TestCounter_Get(t *testing.T) {
	c := New(10)
	if got := c.Get(); got != 10 {
		t.Fatalf("Get() = %d, want 10", got)
	}
}

func TestCounter_Set(t *testing.T) {
	c := New(0)

	for i := 0; i < 1_000; i++ {
		c.Set(i)
	}

	if got := c.Get(); got != 999 {
		t.Fatalf("Get() = %d, want 999", got)
	}
}

func TestCounter_SetThenGet(t *testing.T) {
	c := New(0)

	c.Set(123)
	if got := c.Get(); got != 123 {
		t.Fatalf("Get() = %d, want 123", got)
	}
}
