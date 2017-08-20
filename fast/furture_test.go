package fast

import (
	"testing"
	"fmt"
)

func TestExecQuery(t *testing.T) {
	//初始化Query
	q := query{make(chan string,1),make(chan string,1)}
	//执行Query，注意执行的时候无需准备参数
	ExecQuery(q)
	//准备参数
	q.sql <- "select * froom table"
	//获取结果
	result := <- q.result
	if len(result) != 0 {
		t.Log(result)
	} else {
		t.Fatal(result)
	}
}

func BenchmarkExecQuery(b *testing.B) {
	//初始化Query
	q := query{make(chan string,1),make(chan string,1)}
	for i:=0;i<b.N;i++ {
		//执行Query，注意执行的时候无需准备参数
		ExecQuery(q)
	}

	for n:=0;n<b.N;n++ {
		q.sql <- fmt.Sprintf("select * form %d",n)
		b.Log(<-q.result)
	}
}

func TestMutilExecQuery(t *testing.T) {
	//初始化Query
	q := query{make(chan string,1),make(chan string,1)}
	q2 := query{make(chan string,1),make(chan string,1)}
	//执行Query，注意执行的时候无需准备参数
	data := []query{q,q2}

	MutilExecQuery(data...)

	q.sql <- "1"
	q2.sql <- "2"
	fmt.Println(<-q.result)
	fmt.Println(<-q2.result)
}

func BenchmarkMutilExecQuery(b *testing.B) {
	data := []query{}
	for i:=0;i<b.N;i++ {
		q := query{make(chan string,1),make(chan string,1)}
		q.sql <- fmt.Sprintf("%d",i)
		data = append(data,q)
	}
	MutilExecQuery(data...)
	for _,qq := range data {
		go func() {
			b.Log(<-qq.result)
		}()
	}
}