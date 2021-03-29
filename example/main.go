package main

import (
	"fmt"
	"github.com/lhlyu/future"
	"runtime"
	"time"
)

func task1() int {
	fmt.Println("task1 begining")
	time.Sleep(time.Second)
	fmt.Println("task1 end")
	return 1
}

func task2() int {
	fmt.Println("task2 begining")
	time.Sleep(time.Second * 1)
	fmt.Println("task2 end")
	return 2
}

func task3() int {
	fmt.Println("task3 begining")
	time.Sleep(time.Second * 1)
	fmt.Println("task3 end")
	return 3
}

func main() {
	t1 := future.Async(func() interface{} {
		return task1()
	})
	t2 := future.Async(func() interface{} {
		return task2()
	})
	t3 := future.Async(func() interface{} {
		return task3()
	})
	fmt.Println("miaopas")
	fmt.Println(runtime.NumGoroutine())
	r1 := t1.Await()
	r2 := t2.Await()
	r3 := t3.Await()
	fmt.Println("r1=", r1)
	fmt.Println("r2=", r2)
	fmt.Println("r3=", r3)
}
