package fast

import (
	"fmt"
	"time"
)

type SimpleTask int

func (this *SimpleTask) DoSomeThing() int {
	return <-Rand_generator_3()
}

type Controller int

func (this *Controller) Send() {
	println(time.Now().Format("2006-01-02 15:04:05"))
}

func (this *Controller) ProcessTaskOne() {
	var task SimpleTask = 0
	task.DoSomeThing()
	this.Send()
}

func (this *Controller) ProcessTaskTwo() {
	var task SimpleTask = 0
	go task.DoSomeThing()
	this.Send()
}

func (this *Controller) ProcessTaskThree(inchan chan int, WORKER_NUM int) {
	for i := 0; i < WORKER_NUM; i++ {
		go func() {
			for {
				select {
				case task := <-inchan:
					println(task)
				case <-time.After(time.Second * 1):
					fmt.Println("TimeOut 1s")
				default:
					fmt.Errorf("TASK_CHANNEL is full")
				}
			}
		}()
	}
}
