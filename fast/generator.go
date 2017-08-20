package fast

import (
	"math/rand"
)

//非并发的随机数生成器
func Rand_generator_1() int {
	return rand.Int()
}

//并发的随机数生成器
func Rand_generator_2() chan int {
	//创建通道
	out := make(chan int)
	//创建携程
	go func() {
		 //向通道内写入数据，如果无人读取会等待
		out <- rand.Int()
	}()
	return out
}

//多路复用
func Rand_generator_3() chan int {
	 // 创建两个随机数生成器服务
	rand_generator1 := Rand_generator_2()
	rand_generator2 := Rand_generator_2()
	rand_generator3 := Rand_generator_2()
	rand_generator4 := Rand_generator_2()
	rand_generator5 := Rand_generator_2()
	rand_generator6 := Rand_generator_2()
	rand_generator7 := Rand_generator_2()
	rand_generator8 := Rand_generator_2()
	rand_generator9 := Rand_generator_2()

	//创建通道
	out := make(chan int)

	//创建协程
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator1
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator2
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator3
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator4
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator5
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator6
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator7
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator8
	}()
	go func() {
		//读取生成器1中的数据，整合
		out <- <-rand_generator9
	}()
	return out
}