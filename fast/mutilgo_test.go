package fast

import (
	"testing"
)

func TestController_ProcessTaskOne(t *testing.T) {
	println("TestController_ProcessTaskOne")
	var C Controller = 0
	C.ProcessTaskOne()
}

func BenchmarkController_ProcessTaskOne(b *testing.B) {
	println("BenchmarkController_ProcessTaskOne")
	var C Controller = 0
	for i := 0; i < b.N; i++ {
		C.ProcessTaskOne()
	}
}

func TestController_ProcessTaskTwo(t *testing.T) {
	println("TestController_ProcessTaskTwo")
	var C Controller = 0
	C.ProcessTaskTwo()
}

func BenchmarkController_ProcessTaskTwo(b *testing.B) {
	println("BenchmarkController_ProcessTaskTwo")
	var C Controller = 0
	for i := 0; i < b.N; i++ {
		C.ProcessTaskTwo()
	}
}

func TestController_ProcessTaskThree(t *testing.T) {
	println("TestController_ProcessTaskThree")
	var C Controller = 0
	var TASK_CHANNEL = make(chan int)
	C.ProcessTaskThree(TASK_CHANNEL, 1)
	for i := 0; i < 20; i++ {
		TASK_CHANNEL <- i
	}
}

func BenchmarkController_ProcessTaskThree(b *testing.B) {
	println("TestController_ProcessTaskThree")
	var C Controller = 0
	var TASK_CHANNEL = make(chan int)
	C.ProcessTaskThree(TASK_CHANNEL, 1)
	for i := 0; i < b.N; i++ {
		TASK_CHANNEL <- i
	}
}

func BenchmarkController_ProcessTaskThree2(b *testing.B) {
	println("TestController_ProcessTaskThree")
	var C Controller = 0
	var TASK_CHANNEL = make(chan int)
	C.ProcessTaskThree(TASK_CHANNEL, 1000)
	for i := 0; i < b.N; i++ {
		TASK_CHANNEL <- i
	}
}
