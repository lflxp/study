package fast

import (
	"fmt"
	"math/rand"
)

func doSomething(a,b int) {
	println(a+b)
}

//并发循环
//循环往往是性能上的热点。如果性能瓶颈出现在CPU上的话，那么九成可能性热点是在一个循环体内部。所以如果能让循环体并发执行，那么性能就会提高很多。
func GoForFunc(N int) {
	sem := make(chan int,N)
	data := []int{}
	for s:=0;s<N;s++ {
		data = append(data,rand.Int())
	}
	for i,n := range data {
		go func(a,b int) {
			doSomething(a,b)
		}(i,n)
		sem <- 0
	}
	for z := 0;z<N;z++ {
		fmt.Println(<-sem)
	}
}
