package fast

type query struct {
	//参数Channel
	sql 	chan string
	//结果Channel
	result 	chan string
}

//Furture技术
//Furture是一个很有用的技术，我们常常使用Furture来操作线程。我们可以在使用线程的时候，可以创建一个线程，返回Furture，之后可以通过它等待结果。 但是在协程环境下的Furtue可以更加彻底，输入参数同样可以是Furture的。
//Future是一个非常强大的技术手段。可以在调用的时候不关心数据是否准备好，返回值是否计算好的问题。让程序中的组件在准备好数据的时候自动跑起来。
//执行Query
func ExecQuery(q query) {
	//启动协程
	go func(){
		//获取输入
		sql := <- q.sql
		//访问数据库，输出结果通道
		q.result <- "get "+sql
	}()
	//close(q.sql)
	//close(q.result)
}

//多路复用和furture技术结合
func MutilExecQuery(q... query) {
	for _,qx := range q {
		ExecQuery(qx)
	}
}