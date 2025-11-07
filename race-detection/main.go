package main

type Counter struct {
	n int
}

func New(initial int) *Counter {
	return &Counter{
		n: initial,
	}
}

func (c *Counter) Get() int {
	return c.n
}
