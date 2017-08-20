package fast

import "testing"

func TestFilter(t *testing.T) {
	ch := make(chan int)
	go Generate(ch)

	for i:=0;i<10;i++ {
		prime := <- ch
		println(prime,"n")
		ch1 := make(chan int)
		go Filter(ch,ch1,prime)
		ch = ch1
	}
}
