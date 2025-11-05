package main

import (
	"time"
)

const goRoutinesNum = 20

var (
	mySlice  []int
	myInt    int
	myMap    map[int]int
	myStruct struct{ myInnerInt int }
)

func initData() {
	myInt = 9
	myStruct.myInnerInt = 10
	mySlice = make([]int, goRoutinesNum)
	myMap = make(map[int]int, goRoutinesNum)

	for i := 0; i < goRoutinesNum; i++ {
		mySlice[i] = i
		myMap[i] = i
	}
}

func main() {
	initData()
	timer := time.NewTimer(5 * time.Second)

	for i := 0; i < goRoutinesNum; i++ {
		// int reader
		go func() {
			for {
				_ = myInt
			}
		}()

		// struct reader
		go func() {
			for {
				_ = myStruct.myInnerInt
			}
		}()

		// slice reader
		go func(i int) {
			for {
				_ = mySlice[i]
			}
		}(i)

		// map reader
		go func(i int) {
			for {
				_ = myMap[i]
			}
		}(i)
	}

	<-timer.C
}
