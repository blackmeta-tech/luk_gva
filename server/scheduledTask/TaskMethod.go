package scheduledTask

import (
	"fmt"
	"sync"
	"time"
)

type TaskFunction = func()

func Test() {
	fmt.Println("测试定时任务")
	Testsssss()
	Testsss()
	Testsssss()
}

var TestWd sync.WaitGroup //用户锁
func Testsss() {
	TestWd.Add(1)
	fmt.Println("======Testsss", 1)
	time.Sleep(3 * time.Second)
	TestWd.Done()
}
func Testsssss() {
	TestWd.Add(1)
	fmt.Println("======Testsssss", 2)
	TestWd.Done()
}

// 初始化任务列表
func InitTaskMap() {
	taskMap = map[string]TaskFunction{}
	taskLockMap = map[string]*sync.Mutex{}

	// addTask("test", test)

}
func AddTask(taskName string, task TaskFunction) {
	addTask(taskName, task)
}
