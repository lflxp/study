package fast

import (
	"testing"
	"fmt"
)

func TestSharded_var_whachdog(t *testing.T) {
	//初始化，并开始维护协程
        v := sharded_var{make(chan int), make(chan int)}
        Sharded_var_whachdog(v)

        //读取初始值
        fmt.Println(<-v.reader)
        //写入一个值
        v.writer <- 1
        //读取新写入的值
        fmt.Println(<-v.reader)
}
