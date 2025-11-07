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

func (c *Counter) Set(v int) {
	c.n = v
}
