package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

func exec(cxt context.Context, tasks map[string]func(ctx context.Context) error, c int) error {
	cancelCtx, cancelFnc := context.WithCancel(cxt) // 用于结束正在运行的任务
	var (
		taskErr        unsafe.Pointer           // 记录任务报错
		taskChan       = make(chan struct{}, c) // 控制同时并行的任务
		taskWaiter     = &sync.WaitGroup{}      // 用于等待所有任务完成
		successTaskNum = int32(0)               // 记录成功完成的任务数
		doTaskNum      = int32(0)               // 记录做过的任务数
	)

	for taskName := range tasks {
		select {
		case <-cancelCtx.Done():
			// 跳出任务分配
			break
		default:
			// 标记 chan 才能执行任务
			taskChan <- struct{}{}

			taskFnc := tasks[taskName]
			fmt.Println("------>", taskName, "开始执行", "<--------")
			doTaskNum++

			// do task
			taskWaiter.Add(1)
			go func() {
				defer func() {
					taskWaiter.Done()

					// release chan
					<-taskChan
				}()
				select {
				case <-cancelCtx.Done(): // 任务要提前结束
					return
				default:
					// 执行相关任务
					err := taskFnc(cxt)
					if err == nil {
						atomic.AddInt32(&successTaskNum, 1) // 执行成功的task数量+1
						return
					}

					// 任务执行失败，这个地方可以使用锁，也可以使用原子操作，优先原子操作
					if !atomic.CompareAndSwapPointer(&taskErr, nil, unsafe.Pointer(&err)) {
						return
					}

					cancelFnc() // 结束
				}
			}()
		}
	}

	// 等待任务结束
	taskWaiter.Wait()
	cancelFnc()

	fmt.Println("处理了：", doTaskNum, "个任务")
	fmt.Println("提前结束：", doTaskNum-successTaskNum, "个任务")
	fmt.Println("成功执行了：", successTaskNum, "个任务")
	fmt.Println("剩余：", int32(len(tasks))-doTaskNum, "个任务待处理")

	// 获取错误信息
	return *(*error)(atomic.LoadPointer(&taskErr))
}
