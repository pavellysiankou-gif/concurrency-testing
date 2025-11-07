package main

import "time"

func main() {
	timer := time.NewTimer(3 * time.Second)

	m := map[int]int{1: 1, 2: 2}

	go func() {
		m2 := m
		for {
			_ = m2[1]
			_ = m2[2]
		}
	}()

	go func() {
		for {
			m[1] = 10
			m[2] = 20
		}
	}()

	<-timer.C
}
