package fast

//共享变量有一个读通道和一个写通道组成
type sharded_var struct {
        reader chan int
        writer chan int
}

//一般来说，协程之间不推荐使用共享变量来交互，但是按照这个办法，在一些场合，使用共享变量也是可取的。很多平台上有较为原生的共享变量支持，到底用那种实现比较好，就见仁见智了。
//共享变量维护协程
func Sharded_var_whachdog(v sharded_var) {
        go func() {
                //初始值
                var value int = 0
                for {
                        //监听读写通道，完成服务
                        select {
                        case value = <-v.writer:
                        case v.reader <- value:
                        }
                }
        }()
}
